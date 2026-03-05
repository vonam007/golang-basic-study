package embeddedstructs

import "fmt"

// Base struct
type Address struct {
	Street string
	City   string
}

// FullAddress trả về "Street, City"
func (a Address) FullAddress() string {
	// TODO: Implement this
	_ = fmt.Sprintf("placeholder")
	return ""
}

// Person embed Address
type Person struct {
	Name string
	Age  int
	Address
}

// Greeting trả về "Hi, I'm <Name> from <City>"
func (p Person) Greeting() string {
	// TODO: Implement this
	return ""
}

// Employee embed Person (multi-level embedding)
type Employee struct {
	Person
	Company  string
	Position string
}

// BusinessCard trả về "<Name> - <Position> at <Company>, <City>"
func (e Employee) BusinessCard() string {
	// TODO: Implement this
	return ""
}

// NewEmployee tạo Employee với tất cả thông tin.
func NewEmployee(name string, age int, street, city, company, position string) Employee {
	// TODO: Implement this
	return Employee{}
}
