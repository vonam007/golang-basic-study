package contextvalues

import (
	"context"
	"testing"
)

func TestWithAndGetUserID(t *testing.T) {
	ctx := WithUserID(context.Background(), "user123")
	id, ok := GetUserID(ctx)
	if !ok || id != "user123" {
		t.Errorf("got (%q, %v)", id, ok)
	}
}

func TestGetUserID_Missing(t *testing.T) {
	_, ok := GetUserID(context.Background())
	if ok {
		t.Error("should return false")
	}
}

func TestRequestID(t *testing.T) {
	ctx := WithRequestID(context.Background(), "req-abc")
	id, ok := GetRequestID(ctx)
	if !ok || id != "req-abc" {
		t.Errorf("got (%q, %v)", id, ok)
	}
}

func TestLogLine(t *testing.T) {
	ctx := WithUserID(context.Background(), "alice")
	ctx = WithRequestID(ctx, "req-1")
	got := LogLine(ctx, "hello")
	want := "[req-1] alice: hello"
	if got != want {
		t.Errorf("LogLine = %q; want %q", got, want)
	}

	got2 := LogLine(context.Background(), "test")
	want2 := "[unknown] unknown: test"
	if got2 != want2 {
		t.Errorf("LogLine empty = %q; want %q", got2, want2)
	}
}
