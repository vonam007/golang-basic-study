package bufferedchannels

import (
	"sort"
	"testing"
	"time"
)

func TestProducer(t *testing.T) {
	ch := Producer(5)
	if ch == nil {
		t.Fatal("returned nil")
	}
	var results []int
	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				goto done
			}
			results = append(results, v)
		case <-timeout:
			t.Fatal("timed out")
		}
	}
done:
	if len(results) != 5 {
		t.Errorf("got %d items; want 5", len(results))
	}
}

func TestBatchProcessor(t *testing.T) {
	ch := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	got := BatchProcessor(ch, 2)
	want := [][]int{{0, 1}, {2, 3}, {4}}
	if len(got) != len(want) {
		t.Fatalf("batches = %d; want %d", len(got), len(want))
	}
	for i, batch := range got {
		if len(batch) != len(want[i]) {
			t.Errorf("batch[%d] len = %d", i, len(batch))
		}
	}
}

func TestWorkerPool(t *testing.T) {
	double := func(n int) int { return n * 2 }
	got := WorkerPool([]int{1, 2, 3, 4, 5}, 3, double)
	sort.Ints(got)
	want := []int{2, 4, 6, 8, 10}
	if len(got) != len(want) {
		t.Fatalf("len = %d; want %d", len(got), len(want))
	}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("[%d] = %d; want %d", i, v, want[i])
		}
	}
}
