package customerrors

import (
	"errors"
	"testing"
)

func TestValidationError(t *testing.T) {
	e := &ValidationError{Field: "Name", Message: "cannot be empty"}
	if got := e.Error(); got != "Name: cannot be empty" {
		t.Errorf("Error() = %q", got)
	}
	var _ error = e // must satisfy error interface
}

func TestNotFoundError(t *testing.T) {
	e := &NotFoundError{Resource: "User", ID: "123"}
	if got := e.Error(); got != "User not found: 123" {
		t.Errorf("Error() = %q", got)
	}
}

func TestValidateUser(t *testing.T) {
	if err := ValidateUser("Alice", 30); err != nil {
		t.Errorf("valid user: %v", err)
	}
	err := ValidateUser("", 30)
	if err == nil {
		t.Fatal("empty name should error")
	}
	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Error("should be *ValidationError")
	}

	err = ValidateUser("Bob", -1)
	if err == nil {
		t.Error("negative age should error")
	}

	err = ValidateUser("Bob", 151)
	if err == nil {
		t.Error("age > 150 should error")
	}
}

func TestFindUser(t *testing.T) {
	users := map[string]string{"1": "Alice", "2": "Bob"}
	name, err := FindUser(users, "1")
	if err != nil || name != "Alice" {
		t.Errorf("FindUser(1) = (%q, %v)", name, err)
	}

	_, err = FindUser(users, "99")
	if err == nil {
		t.Fatal("missing user should error")
	}
	var nfe *NotFoundError
	if !errors.As(err, &nfe) {
		t.Error("should be *NotFoundError")
	}
	if nfe.ID != "99" {
		t.Errorf("NotFoundError.ID = %q; want 99", nfe.ID)
	}
}
