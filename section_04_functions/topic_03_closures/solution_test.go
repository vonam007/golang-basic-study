package closures

import "testing"

func TestCounter(t *testing.T) {
	c := Counter()
	if c == nil {
		t.Fatal("Counter() returned nil")
	}
	for i := 1; i <= 5; i++ {
		if got := c(); got != i {
			t.Errorf("Counter call %d = %d; want %d", i, got, i)
		}
	}
	c2 := Counter()
	if got := c2(); got != 1 {
		t.Error("New Counter should start from 1")
	}
}

func TestAdder(t *testing.T) {
	add5 := Adder(5)
	if add5 == nil {
		t.Fatal("Adder returned nil")
	}
	if got := add5(3); got != 8 {
		t.Errorf("Adder(5)(3) = %d; want 8", got)
	}
	if got := add5(0); got != 5 {
		t.Errorf("Adder(5)(0) = %d; want 5", got)
	}
	if got := add5(-5); got != 0 {
		t.Errorf("Adder(5)(-5) = %d; want 0", got)
	}

	add0 := Adder(0)
	if got := add0(7); got != 7 {
		t.Errorf("Adder(0)(7) = %d; want 7", got)
	}
}

func TestFibonacci(t *testing.T) {
	fib := Fibonacci()
	if fib == nil {
		t.Fatal("Fibonacci returned nil")
	}
	want := []int{0, 1, 1, 2, 3, 5, 8, 13}
	for i, w := range want {
		if got := fib(); got != w {
			t.Errorf("Fibonacci call %d = %d; want %d", i, got, w)
		}
	}
}

func TestAccumulator(t *testing.T) {
	acc := Accumulator()
	if acc == nil {
		t.Fatal("Accumulator returned nil")
	}
	if got := acc(5); got != 5 {
		t.Errorf("acc(5) = %d; want 5", got)
	}
	if got := acc(3); got != 8 {
		t.Errorf("acc(3) = %d; want 8", got)
	}
	if got := acc(-2); got != 6 {
		t.Errorf("acc(-2) = %d; want 6", got)
	}
	if got := acc(0); got != 6 {
		t.Errorf("acc(0) = %d; want 6", got)
	}
}
