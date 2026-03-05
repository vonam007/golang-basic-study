package genericds

import (
	"errors"
	"testing"
)

func TestStack(t *testing.T) {
	s := &Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Len() != 3 {
		t.Errorf("Len = %d", s.Len())
	}
	v, ok := s.Pop()
	if !ok || v != 3 {
		t.Errorf("Pop = (%d, %v)", v, ok)
	}
	v, ok = s.Peek()
	if !ok || v != 2 {
		t.Errorf("Peek = (%d, %v)", v, ok)
	}
	s.Pop()
	s.Pop()
	_, ok = s.Pop()
	if ok {
		t.Error("Pop empty should fail")
	}
}

func TestStackString(t *testing.T) {
	s := &Stack[string]{}
	s.Push("hello")
	v, ok := s.Pop()
	if !ok || v != "hello" {
		t.Errorf("Pop = (%q, %v)", v, ok)
	}
}

func TestResult(t *testing.T) {
	ok := NewSuccess(42)
	if !ok.IsOk() {
		t.Error("should be ok")
	}
	if ok.Unwrap() != 42 {
		t.Error("Unwrap")
	}

	fail := NewFailure[int](errors.New("fail"))
	if fail.IsOk() {
		t.Error("should not be ok")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Unwrap failure should panic")
		}
	}()
	fail.Unwrap()
}
