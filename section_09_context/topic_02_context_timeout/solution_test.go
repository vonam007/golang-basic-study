package contexttimeout

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestFetchWithTimeout_Success(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	got, err := FetchWithTimeout(ctx, 50*time.Millisecond)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if got == "" {
		t.Error("should return result")
	}
}

func TestFetchWithTimeout_Timeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	_, err := FetchWithTimeout(ctx, time.Second)
	if err == nil {
		t.Error("should timeout")
	}
}

func TestRetryWithTimeout(t *testing.T) {
	attempts := 0
	op := func() error {
		attempts++
		if attempts < 3 {
			return errors.New("fail")
		}
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	err := RetryWithTimeout(ctx, op, 5)
	if err != nil {
		t.Errorf("should succeed after retries: %v", err)
	}
}
