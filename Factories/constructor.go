package Factories

import (
	"errors"
)

type Person interface {
	Name() string
	Age() int
}

// Immutable
type person struct {
	name string
	age  int
}

func (p *person) Name() string {
	return p.name
}

func (p *person) Age() int {
	return p.age
}

func NewPerson(name string, age int) (Person, error) {
	if len(name) == 0 {
		return nil, errors.New("name cannot be blank")
	}
	if age < 0 {
		return nil, errors.New("age cannot be negative")
	}
	return &person{name, age}, nil
}
