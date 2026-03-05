package errorwrapping

import (
	"errors"
	"testing"
)

func TestWrapError(t *testing.T) {
	original := errors.New("disk full")
	wrapped := WrapError("save file", original)
	if wrapped == nil {
		t.Fatal("WrapError returned nil")
	}
	if !errors.Is(wrapped, original) {
		t.Error("wrapped error should contain original")
	}
}

func TestIsNotFound(t *testing.T) {
	if IsNotFound(nil) {
		t.Error("nil should not be not-found")
	}
	if !IsNotFound(ErrNotFound) {
		t.Error("ErrNotFound should match")
	}
	wrapped := WrapError("get user", ErrNotFound)
	if !IsNotFound(wrapped) {
		t.Error("wrapped ErrNotFound should match")
	}
	if IsNotFound(ErrUnauthorized) {
		t.Error("ErrUnauthorized should not match")
	}
}

func TestGetUser(t *testing.T) {
	name, err := GetUser("1")
	if err != nil || name == "" {
		t.Errorf("GetUser(1) = (%q, %v)", name, err)
	}

	_, err = GetUser("999")
	if err == nil {
		t.Fatal("GetUser(999) should error")
	}
	if !errors.Is(err, ErrNotFound) {
		t.Error("GetUser error should wrap ErrNotFound")
	}
}

func TestGetUserProfile(t *testing.T) {
	_, err := GetUserProfile("999")
	if err == nil {
		t.Fatal("should error")
	}
	if !errors.Is(err, ErrNotFound) {
		t.Error("should still contain ErrNotFound in chain")
	}
}
