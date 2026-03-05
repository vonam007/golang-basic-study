package contextbasics

import (
	"context"
	"testing"
	"time"
)

func TestDoWorkWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() { time.Sleep(100 * time.Millisecond); cancel() }()
	count := DoWorkWithCancel(ctx)
	if count == 0 {
		t.Error("should have done some work")
	}
}

func TestPropagateContext(t *testing.T) {
	ctx := context.Background()
	got, err := PropagateContext(ctx)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if got == "" {
		t.Error("should return result")
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err = PropagateContext(ctx)
	if err == nil {
		t.Error("cancelled context should error")
	}
}

func TestCancelAfterN(t *testing.T) {
	items := []string{"a", "b", "c", "d", "e"}
	got := CancelAfterN(items, 3)
	if len(got) != 3 {
		t.Errorf("got %d items; want 3", len(got))
	}
	if got[0] != "a" || got[1] != "b" || got[2] != "c" {
		t.Errorf("got %v; want [a b c]", got)
	}
}
