package panicrecover

import "testing"

func TestSafeExecute_NoPanic(t *testing.T) {
	err := SafeExecute(func() { _ = 1 + 1 })
	if err != nil {
		t.Errorf("no panic: err = %v; want nil", err)
	}
}

func TestSafeExecute_WithPanic(t *testing.T) {
	err := SafeExecute(func() { panic("boom") })
	if err == nil {
		t.Fatal("should catch panic")
	}
}

func TestMustPositive(t *testing.T) {
	if got := MustPositive(5); got != 5 {
		t.Errorf("MustPositive(5) = %d", got)
	}
	if got := MustPositive(0); got != 0 {
		t.Errorf("MustPositive(0) = %d", got)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("MustPositive(-1) should panic")
		}
	}()
	MustPositive(-1)
}

func TestSafePositive(t *testing.T) {
	got, err := SafePositive(5)
	if err != nil || got != 5 {
		t.Errorf("SafePositive(5) = (%d, %v)", got, err)
	}

	got, err = SafePositive(-1)
	if err == nil {
		t.Error("SafePositive(-1) should return error")
	}
	if got != 0 {
		t.Errorf("SafePositive(-1) value = %d; want 0", got)
	}
}

func TestRecoverMiddleware(t *testing.T) {
	handler := func(s string) string {
		if s == "" {
			panic("empty input")
		}
		return "ok: " + s
	}
	safe := RecoverMiddleware(handler)
	if safe == nil {
		t.Fatal("RecoverMiddleware returned nil")
	}

	got, err := safe("test")
	if err != nil || got != "ok: test" {
		t.Errorf("safe(test) = (%q, %v)", got, err)
	}

	_, err = safe("")
	if err == nil {
		t.Error("safe('') should return error from recovered panic")
	}
}
