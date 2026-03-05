package waitgroup

import (
	"errors"
	"testing"
)

func TestConcurrentFetch(t *testing.T) {
	fetcher := func(url string) string { return "data:" + url }
	urls := []string{"a.com", "b.com", "c.com"}
	got := ConcurrentFetch(urls, fetcher)
	if got == nil {
		t.Fatal("returned nil")
	}
	for _, u := range urls {
		if got[u] != "data:"+u {
			t.Errorf("[%s] = %q", u, got[u])
		}
	}
}

func TestParallelProcess(t *testing.T) {
	double := func(n int) int { return n * 2 }
	got := ParallelProcess([]int{1, 2, 3}, double)
	want := []int{2, 4, 6}
	if len(got) != len(want) {
		t.Fatalf("len = %d", len(got))
	}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("[%d] = %d; want %d", i, v, want[i])
		}
	}
}

func TestErrGroup_AllOK(t *testing.T) {
	tasks := []func() error{
		func() error { return nil },
		func() error { return nil },
	}
	if err := ErrGroup(tasks); err != nil {
		t.Errorf("all ok: %v", err)
	}
}

func TestErrGroup_WithError(t *testing.T) {
	tasks := []func() error{
		func() error { return nil },
		func() error { return errors.New("fail") },
	}
	if err := ErrGroup(tasks); err == nil {
		t.Error("should return error")
	}
}
