package deferexample

import "testing"

func TestDeferOrder(t *testing.T) {
	got := DeferOrder()
	want := []string{"3", "2", "1"}
	if len(got) != len(want) {
		t.Fatalf("DeferOrder() len = %d; want %d", len(got), len(want))
	}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("DeferOrder()[%d] = %q; want %q", i, v, want[i])
		}
	}
}

func TestSafeDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr bool
	}{
		{"normal", 10, 2, 5, false},
		{"zero div", 10, 0, 0, true},
		{"negative", -10, 2, -5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SafeDivide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("SafeDivide(%d, %d) err = %v; wantErr %v", tt.a, tt.b, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SafeDivide(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestStackTrace(t *testing.T) {
	got := StackTrace([]string{"main", "helper"})
	want := []string{"enter main", "enter helper", "exit helper", "exit main"}
	if len(got) != len(want) {
		t.Fatalf("StackTrace len = %d; want %d\ngot: %v", len(got), len(want), got)
	}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("StackTrace()[%d] = %q; want %q", i, v, want[i])
		}
	}
}

func TestStackTrace_Single(t *testing.T) {
	got := StackTrace([]string{"A"})
	want := []string{"enter A", "exit A"}
	if len(got) != len(want) {
		t.Fatalf("StackTrace len = %d; want %d", len(got), len(want))
	}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("StackTrace()[%d] = %q; want %q", i, v, want[i])
		}
	}
}

func TestStackTrace_Empty(t *testing.T) {
	got := StackTrace([]string{})
	if len(got) != 0 {
		t.Errorf("StackTrace([]) = %v; want empty", got)
	}
}

func TestCountDown(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{5, []int{5, 4, 3, 2, 1}},
		{1, []int{1}},
		{0, nil},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := CountDown(tt.n)
			if len(got) != len(tt.want) {
				t.Fatalf("CountDown(%d) len = %d; want %d", tt.n, len(got), len(tt.want))
			}
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("CountDown(%d)[%d] = %d; want %d", tt.n, i, v, tt.want[i])
				}
			}
		})
	}
}
