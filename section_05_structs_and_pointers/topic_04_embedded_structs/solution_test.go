package embeddedstructs

import "testing"

func TestFullAddress(t *testing.T) {
	a := Address{Street: "123 Main St", City: "Hanoi"}
	want := "123 Main St, Hanoi"
	if got := a.FullAddress(); got != want {
		t.Errorf("FullAddress() = %q; want %q", got, want)
	}
}

func TestGreeting(t *testing.T) {
	p := Person{Name: "Alice", Age: 30, Address: Address{City: "Hanoi"}}
	want := "Hi, I'm Alice from Hanoi"
	if got := p.Greeting(); got != want {
		t.Errorf("Greeting() = %q; want %q", got, want)
	}
}

func TestPromotedFields(t *testing.T) {
	p := Person{Name: "Bob", Address: Address{City: "HCMC", Street: "456 Le Loi"}}
	if p.City != "HCMC" {
		t.Errorf("Promoted City = %q; want HCMC", p.City)
	}
	if p.Street != "456 Le Loi" {
		t.Errorf("Promoted Street = %q; want 456 Le Loi", p.Street)
	}
}

func TestBusinessCard(t *testing.T) {
	e := Employee{
		Person:   Person{Name: "Charlie", Age: 28, Address: Address{City: "Danang"}},
		Company:  "Google",
		Position: "Engineer",
	}
	want := "Charlie - Engineer at Google, Danang"
	if got := e.BusinessCard(); got != want {
		t.Errorf("BusinessCard() = %q; want %q", got, want)
	}
}

func TestNewEmployee(t *testing.T) {
	e := NewEmployee("Alice", 30, "123 Main", "Hanoi", "Google", "SWE")
	if e.Name != "Alice" || e.Age != 30 || e.City != "Hanoi" || e.Company != "Google" {
		t.Errorf("NewEmployee = %+v", e)
	}
}

func TestMultiLevelAccess(t *testing.T) {
	e := NewEmployee("Bob", 25, "456 St", "HCMC", "Meta", "PM")
	// Should access through multiple levels
	if e.City != "HCMC" {
		t.Errorf("e.City = %q; want HCMC", e.City)
	}
	if e.Name != "Bob" {
		t.Errorf("e.Name = %q; want Bob", e.Name)
	}
}
