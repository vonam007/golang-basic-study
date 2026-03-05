package basicerrors

import "testing"

func TestDivide(t *testing.T) {
	got, err := Divide(10, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got < 3.33 || got > 3.34 {
		t.Errorf("Divide(10,3) = %v", got)
	}
	_, err = Divide(1, 0)
	if err == nil {
		t.Error("Divide(1,0) should error")
	}
}

func TestParseAge(t *testing.T) {
	tests := []struct {
		s       string
		want    int
		wantErr bool
	}{
		{"25", 25, false}, {"0", 0, false}, {"150", 150, false},
		{"-1", 0, true}, {"151", 0, true}, {"abc", 0, true}, {"", 0, true},
	}
	for _, tt := range tests {
		got, err := ParseAge(tt.s)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseAge(%q) err = %v; wantErr %v", tt.s, err, tt.wantErr)
			continue
		}
		if !tt.wantErr && got != tt.want {
			t.Errorf("ParseAge(%q) = %d; want %d", tt.s, got, tt.want)
		}
	}
}

func TestSafeGet(t *testing.T) {
	items := []string{"a", "b", "c"}
	got, err := SafeGet(items, 1)
	if err != nil || got != "b" {
		t.Errorf("SafeGet(1) = (%q, %v)", got, err)
	}
	_, err = SafeGet(items, 5)
	if err == nil {
		t.Error("SafeGet(5) should error")
	}
	_, err = SafeGet(items, -1)
	if err == nil {
		t.Error("SafeGet(-1) should error")
	}
}

func TestMultiError(t *testing.T) {
	ops := []func() error{
		func() error { return nil },
		func() error { return nil },
	}
	if err := MultiError(ops); err != nil {
		t.Errorf("MultiError(all ok) = %v; want nil", err)
	}

	errOps := []func() error{
		func() error { return nil },
		func() error { _, e := Divide(1, 0); return e },
	}
	if err := MultiError(errOps); err == nil {
		t.Error("MultiError should return error")
	}
}
