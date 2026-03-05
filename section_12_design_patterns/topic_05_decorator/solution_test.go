package decorator

import "testing"

func identity(s string) string { return s }

func TestWithPrefix(t *testing.T) {
	fn := WithPrefix("[LOG] ")(identity)
	if got := fn("hello"); got != "[LOG] hello" {
		t.Errorf("got %q", got)
	}
}

func TestWithSuffix(t *testing.T) {
	fn := WithSuffix("!")(identity)
	if got := fn("hello"); got != "hello!" {
		t.Errorf("got %q", got)
	}
}

func TestWithUpperCase(t *testing.T) {
	fn := WithUpperCase()(identity)
	if got := fn("hello"); got != "HELLO" {
		t.Errorf("got %q", got)
	}
}

func TestChain(t *testing.T) {
	fn := Chain(identity, WithPrefix("["), WithSuffix("]"), WithUpperCase())
	got := fn("hello")
	want := "[HELLO]"
	if got != want {
		t.Errorf("Chain = %q; want %q", got, want)
	}
}
