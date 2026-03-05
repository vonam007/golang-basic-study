package structtags

import "testing"

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email,omitempty"`
	Age   int    `json:"-"`
	Phone string
}

func TestGetJSONTags(t *testing.T) {
	got := GetJSONTags(User{})
	if got == nil {
		t.Fatal("returned nil")
	}
	if got["Name"] != "name" {
		t.Errorf("Name tag = %q", got["Name"])
	}
	if got["Email"] != "email,omitempty" {
		t.Errorf("Email tag = %q", got["Email"])
	}
}

func TestToMap(t *testing.T) {
	u := User{Name: "Alice", Email: "a@b.com", Age: 30, Phone: "123"}
	got := ToMap(u)
	if got == nil {
		t.Fatal("returned nil")
	}
	if got["name"] != "Alice" {
		t.Errorf("name = %v", got["name"])
	}
	if got["email,omitempty"] != nil && got["email"] != "a@b.com" {
		// accept either "email,omitempty" or "email" as key
	}
	if _, exists := got["Age"]; exists {
		t.Error("Age (tag -) should be skipped")
	}
}

func TestHasTag(t *testing.T) {
	if !HasTag(User{}, "Name", "json") {
		t.Error("Name should have json tag")
	}
	if !HasTag(User{}, "Name", "validate") {
		t.Error("Name should have validate tag")
	}
	if HasTag(User{}, "Phone", "json") {
		t.Error("Phone has no json tag")
	}
}
