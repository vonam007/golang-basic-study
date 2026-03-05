package selecttopic

import (
	"testing"
	"time"
)

func TestFirstResponse(t *testing.T) {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch1 <- "fast"
	got, ok := FirstResponse(ch1, ch2, time.Second)
	if !ok || got != "fast" {
		t.Errorf("got (%q, %v); want (fast, true)", got, ok)
	}
}

func TestFirstResponse_Timeout(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	_, ok := FirstResponse(ch1, ch2, 50*time.Millisecond)
	if ok {
		t.Error("should timeout")
	}
}

func TestMerge(t *testing.T) {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	close(ch1)
	ch2 <- 3
	ch2 <- 4
	close(ch2)
	got := Merge(ch1, ch2)
	if len(got) != 4 {
		t.Errorf("Merge got %d items; want 4", len(got))
	}
}

func TestTicker(t *testing.T) {
	done := make(chan struct{})
	go func() { time.Sleep(250 * time.Millisecond); close(done) }()
	count := Ticker(50*time.Millisecond, done)
	if count < 3 || count > 6 {
		t.Errorf("Ticker count = %d; expected 3-6", count)
	}
}
