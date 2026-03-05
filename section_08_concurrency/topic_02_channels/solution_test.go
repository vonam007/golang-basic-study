package channels

import (
	"sort"
	"testing"
	"time"
)

// ============================================================
// Tests for Sum
// ============================================================

// helper: receive from channel with timeout, fatal on deadlock
func recvIntWithTimeout(t *testing.T, ch <-chan int, label string) int {
	t.Helper()
	select {
	case val := <-ch:
		t.Logf("[%s] ✅ received value: %d", label, val)
		return val
	case <-time.After(2 * time.Second):
		t.Fatalf("[%s] ❌ timed out waiting for channel - Sum did not send a value (deadlock)", label)
		return 0
	}
}

func TestSum_BasicCase(t *testing.T) {
	t.Log("▶ Testing Sum([1,2,3,4,5]) — expect 15")
	ch := make(chan int)
	go Sum([]int{1, 2, 3, 4, 5}, ch)

	got := recvIntWithTimeout(t, ch, "BasicCase")
	want := 15
	if got != want {
		t.Errorf("Sum([1,2,3,4,5]) = %d; want %d", got, want)
	}
}

func TestSum_EmptySlice(t *testing.T) {
	t.Log("▶ Testing Sum([]) — expect 0")
	ch := make(chan int)
	go Sum([]int{}, ch)

	got := recvIntWithTimeout(t, ch, "EmptySlice")
	want := 0
	if got != want {
		t.Errorf("Sum([]) = %d; want %d", got, want)
	}
}

func TestSum_SingleElement(t *testing.T) {
	t.Log("▶ Testing Sum([42]) — expect 42")
	ch := make(chan int)
	go Sum([]int{42}, ch)

	got := recvIntWithTimeout(t, ch, "SingleElement")
	want := 42
	if got != want {
		t.Errorf("Sum([42]) = %d; want %d", got, want)
	}
}

func TestSum_NegativeNumbers(t *testing.T) {
	t.Log("▶ Testing Sum([-1,-2,-3]) — expect -6")
	ch := make(chan int)
	go Sum([]int{-1, -2, -3}, ch)

	got := recvIntWithTimeout(t, ch, "NegativeNumbers")
	want := -6
	if got != want {
		t.Errorf("Sum([-1,-2,-3]) = %d; want %d", got, want)
	}
}

func TestSum_MixedNumbers(t *testing.T) {
	t.Log("▶ Testing Sum([10,-5,3,-8,20]) — expect 20")
	ch := make(chan int)
	go Sum([]int{10, -5, 3, -8, 20}, ch)

	got := recvIntWithTimeout(t, ch, "MixedNumbers")
	want := 20
	if got != want {
		t.Errorf("Sum([10,-5,3,-8,20]) = %d; want %d", got, want)
	}
}

func TestSum_ConcurrentSplit(t *testing.T) {
	t.Log("▶ Testing concurrent Sum split — expect 55")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mid := len(nums) / 2

	ch := make(chan int, 2)
	go Sum(nums[:mid], ch)
	go Sum(nums[mid:], ch)

	first := recvIntWithTimeout(t, ch, "ConcurrentSplit-1")
	second := recvIntWithTimeout(t, ch, "ConcurrentSplit-2")
	total := first + second
	want := 55
	if total != want {
		t.Errorf("Concurrent Sum = %d; want %d", total, want)
	}
}

// ============================================================
// Tests for FanIn
// ============================================================

func TestFanIn_BasicMerge(t *testing.T) {
	ch1 := make(chan string, 3)
	ch2 := make(chan string, 3)

	ch1 <- "a"
	ch1 <- "b"
	ch1 <- "c"
	close(ch1)

	ch2 <- "x"
	ch2 <- "y"
	close(ch2)

	merged := FanIn(ch1, ch2)
	if merged == nil {
		t.Fatal("FanIn returned nil channel")
	}

	var results []string
	timeout := time.After(2 * time.Second)

	for i := 0; i < 5; i++ {
		select {
		case val, ok := <-merged:
			if !ok {
				// channel closed early, collect what we have
				goto done
			}
			results = append(results, val)
		case <-timeout:
			t.Fatal("FanIn timed out - possible deadlock")
		}
	}
done:

	if len(results) != 5 {
		t.Errorf("FanIn got %d values; want 5", len(results))
	}

	sort.Strings(results)
	want := []string{"a", "b", "c", "x", "y"}
	for i, v := range results {
		if v != want[i] {
			t.Errorf("results[%d] = %q; want %q", i, v, want[i])
		}
	}
}

func TestFanIn_EmptyChannels(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	close(ch1)
	close(ch2)

	merged := FanIn(ch1, ch2)
	if merged == nil {
		t.Fatal("FanIn returned nil channel")
	}

	timeout := time.After(2 * time.Second)
	select {
	case _, ok := <-merged:
		if ok {
			t.Error("Expected closed channel, but received a value")
		}
	case <-timeout:
		t.Fatal("FanIn did not close merged channel - possible deadlock")
	}
}

func TestFanIn_OneEmptyChannel(t *testing.T) {
	ch1 := make(chan string, 2)
	ch2 := make(chan string)

	ch1 <- "hello"
	ch1 <- "world"
	close(ch1)
	close(ch2)

	merged := FanIn(ch1, ch2)
	if merged == nil {
		t.Fatal("FanIn returned nil channel")
	}

	var results []string
	timeout := time.After(2 * time.Second)

	for {
		select {
		case val, ok := <-merged:
			if !ok {
				goto done
			}
			results = append(results, val)
		case <-timeout:
			t.Fatal("FanIn timed out")
		}
	}
done:

	if len(results) != 2 {
		t.Errorf("FanIn got %d values; want 2", len(results))
	}
}

func TestFanIn_ClosesWhenBothInputsClosed(t *testing.T) {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	ch1 <- "one"
	close(ch1)

	ch2 <- "two"
	close(ch2)

	merged := FanIn(ch1, ch2)
	if merged == nil {
		t.Fatal("FanIn returned nil channel")
	}

	count := 0
	timeout := time.After(2 * time.Second)

	for {
		select {
		case _, ok := <-merged:
			if !ok {
				if count != 2 {
					t.Errorf("Channel closed after %d values; want 2", count)
				}
				return
			}
			count++
		case <-timeout:
			t.Fatal("Merged channel was not closed - possible goroutine leak")
		}
	}
}

// ============================================================
// Tests for Pipeline
// ============================================================

func TestPipeline_BasicCase(t *testing.T) {
	// nums: [1, 2, 3] → square: [1, 4, 9] → filter (>5): [9]
	out := Pipeline([]int{1, 2, 3}, 5)
	if out == nil {
		t.Fatal("Pipeline returned nil channel")
	}

	var results []int
	timeout := time.After(2 * time.Second)

	for {
		select {
		case val, ok := <-out:
			if !ok {
				goto done
			}
			results = append(results, val)
		case <-timeout:
			t.Fatal("Pipeline timed out")
		}
	}
done:

	want := []int{9}
	if len(results) != len(want) {
		t.Fatalf("Pipeline got %d values; want %d", len(results), len(want))
	}
	for i, v := range results {
		if v != want[i] {
			t.Errorf("results[%d] = %d; want %d", i, v, want[i])
		}
	}
}

func TestPipeline_AllPass(t *testing.T) {
	// nums: [5, 6] → square: [25, 36] → filter (>0): [25, 36]
	out := Pipeline([]int{5, 6}, 0)
	if out == nil {
		t.Fatal("Pipeline returned nil channel")
	}

	var results []int
	timeout := time.After(2 * time.Second)

	for {
		select {
		case val, ok := <-out:
			if !ok {
				goto done
			}
			results = append(results, val)
		case <-timeout:
			t.Fatal("Pipeline timed out")
		}
	}
done:

	sort.Ints(results)
	want := []int{25, 36}
	if len(results) != len(want) {
		t.Fatalf("Pipeline got %d values; want %d", len(results), len(want))
	}
	for i, v := range results {
		if v != want[i] {
			t.Errorf("results[%d] = %d; want %d", i, v, want[i])
		}
	}
}

func TestPipeline_NonePass(t *testing.T) {
	// nums: [1, 2] → square: [1, 4] → filter (>100): []
	out := Pipeline([]int{1, 2}, 100)
	if out == nil {
		t.Fatal("Pipeline returned nil channel")
	}

	timeout := time.After(2 * time.Second)
	select {
	case val, ok := <-out:
		if ok {
			t.Errorf("Expected no values, but got %d", val)
		}
	case <-timeout:
		t.Fatal("Pipeline did not close channel")
	}
}

func TestPipeline_EmptyInput(t *testing.T) {
	out := Pipeline([]int{}, 0)
	if out == nil {
		t.Fatal("Pipeline returned nil channel")
	}

	timeout := time.After(2 * time.Second)
	select {
	case _, ok := <-out:
		if ok {
			t.Error("Expected closed channel for empty input")
		}
	case <-timeout:
		t.Fatal("Pipeline did not close channel for empty input")
	}
}

func TestPipeline_NegativeNumbers(t *testing.T) {
	// nums: [-3, -2, 4] → square: [9, 4, 16] → filter (>5): [9, 16]
	out := Pipeline([]int{-3, -2, 4}, 5)
	if out == nil {
		t.Fatal("Pipeline returned nil channel")
	}

	var results []int
	timeout := time.After(2 * time.Second)

	for {
		select {
		case val, ok := <-out:
			if !ok {
				goto done
			}
			results = append(results, val)
		case <-timeout:
			t.Fatal("Pipeline timed out")
		}
	}
done:

	sort.Ints(results)
	want := []int{9, 16}
	if len(results) != len(want) {
		t.Fatalf("Pipeline got %d values; want %d", len(results), len(want))
	}
	for i, v := range results {
		if v != want[i] {
			t.Errorf("results[%d] = %d; want %d", i, v, want[i])
		}
	}
}

func TestPipeline_LargeInput(t *testing.T) {
	// 100 elements, only elements where x^2 > 9000 pass
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i + 1
	}

	out := Pipeline(nums, 9000)
	if out == nil {
		t.Fatal("Pipeline returned nil channel")
	}

	var results []int
	timeout := time.After(5 * time.Second)

	for {
		select {
		case val, ok := <-out:
			if !ok {
				goto done
			}
			results = append(results, val)
		case <-timeout:
			t.Fatal("Pipeline timed out on large input")
		}
	}
done:

	// x^2 > 9000 when x >= 95 (95^2 = 9025)
	// So x = 95, 96, 97, 98, 99, 100 → 6 values
	if len(results) != 6 {
		t.Errorf("Large Pipeline got %d values; want 6", len(results))
	}

	for _, v := range results {
		if v <= 9000 {
			t.Errorf("Pipeline output %d should be > 9000", v)
		}
	}
}
