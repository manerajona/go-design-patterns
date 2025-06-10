package prototype

import (
	"bytes"
	"encoding/gob"
)

type Location struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Location
}

func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	d := gob.NewDecoder(&b)
	dCopy := Employee{}
	_ = d.Decode(&dCopy)
	return &dCopy
}

// Prototype Factories
var londonOffice = Employee{
	"",
	Location{0, "123 East Dr", "London"},
}
var parisOffice = Employee{
	"",
	Location{0, "66 Victor Hugo Ave", "Paris"},
}

// Constructors
func newEmployee(prototype *Employee, name string, suite int) (employee *Employee) {
	employee = prototype.DeepCopy()
	employee.Name = name
	employee.Office.Suite = suite
	return
}

func NewLondonOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&londonOffice, name, suite)
}

func NewParisOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&parisOffice, name, suite)
}
