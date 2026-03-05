package dynamicinvocation

import "testing"

type Calculator struct{ Value int }

func (c Calculator) Double() int   { return c.Value * 2 }
func (c Calculator) Add(n int) int { return c.Value + n }

func TestCallMethod(t *testing.T) {
	c := Calculator{Value: 5}
	results, err := CallMethod(c, "Double")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if len(results) != 1 || results[0] != 10 {
		t.Errorf("Double() = %v; want [10]", results)
	}
	_, err = CallMethod(c, "Missing")
	if err == nil {
		t.Error("missing method should error")
	}
}

func TestSetField(t *testing.T) {
	c := &Calculator{Value: 5}
	err := SetField(c, "Value", 10)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if c.Value != 10 {
		t.Errorf("Value = %d; want 10", c.Value)
	}
	err = SetField(c, "Missing", 1)
	if err == nil {
		t.Error("missing field should error")
	}
}

func TestMethodNames(t *testing.T) {
	c := Calculator{}
	names := MethodNames(c)
	if len(names) < 2 {
		t.Errorf("should have >= 2 methods, got %v", names)
	}
}
