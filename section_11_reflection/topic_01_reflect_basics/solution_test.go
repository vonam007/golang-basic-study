package reflectbasics

import (
	"reflect"
	"testing"
)

type Sample struct {
	Name string
	Age  int
}

func TestTypeInfo(t *testing.T) {
	name, kind := TypeInfo(42)
	if name != "int" || kind != reflect.Int {
		t.Errorf("int: (%q, %v)", name, kind)
	}
	name, kind = TypeInfo("hello")
	if name != "string" || kind != reflect.String {
		t.Errorf("string: (%q, %v)", name, kind)
	}
	name, kind = TypeInfo(Sample{})
	if kind != reflect.Struct {
		t.Errorf("struct: (%q, %v)", name, kind)
	}
}

func TestFieldNames(t *testing.T) {
	got := FieldNames(Sample{Name: "Go", Age: 15})
	want := []string{"Name", "Age"}
	if len(got) != len(want) {
		t.Fatalf("len = %d", len(got))
	}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("[%d] = %q; want %q", i, v, want[i])
		}
	}
	if FieldNames(42) != nil {
		t.Error("non-struct should return nil")
	}
}

func TestGetField(t *testing.T) {
	s := Sample{Name: "Go", Age: 15}
	v, ok := GetField(s, "Name")
	if !ok || v != "Go" {
		t.Errorf("GetField Name = (%v, %v)", v, ok)
	}
	_, ok = GetField(s, "Missing")
	if ok {
		t.Error("missing field should fail")
	}
}

func TestIsZero(t *testing.T) {
	if !IsZero(0) {
		t.Error("0 is zero")
	}
	if !IsZero("") {
		t.Error("empty string is zero")
	}
	if IsZero(42) {
		t.Error("42 is not zero")
	}
	if !IsZero(false) {
		t.Error("false is zero")
	}
}
