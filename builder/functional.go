package builder

import (
	"regexp"
)

var (
	fullNameRegEx = regexp.MustCompile(`^[a-zA-Z]+([-' ]?[a-zA-Z]+)*$`)
	phoneRegEx    = regexp.MustCompile(`^\+?[0-9\s\-().]{7,20}$`)
)

type Employee struct {
	name, phone string
	salary      int
}

type EmployeeConsumer func(employee *Employee)

type EmployeeBuilder struct {
	consumers []EmployeeConsumer
}

func (builder *EmployeeBuilder) Name(name string) *EmployeeBuilder {
	consumer := func(employee *Employee) {
		if !fullNameRegEx.MatchString(name) {
			panic("Full Name incorrect")
		}
		employee.name = name
	}
	builder.consumers = append(builder.consumers, consumer)
	return builder
}

func (builder *EmployeeBuilder) Phone(phone string) *EmployeeBuilder {
	consumer := func(p *Employee) {
		if !phoneRegEx.MatchString(phone) {
			panic("Phone incorrect")
		}
		p.phone = phone
	}
	builder.consumers = append(builder.consumers, consumer)
	return builder
}

func (builder *EmployeeBuilder) Salary(salary int) *EmployeeBuilder {
	consumer := func(p *Employee) {
		if salary <= 25_000 || salary >= 1_000_000 {
			panic("Salary incorrect")
		}
		p.salary = salary
	}
	builder.consumers = append(builder.consumers, consumer)
	return builder
}

func (builder *EmployeeBuilder) Build() *Employee {
	employee := &Employee{}
	for _, consume := range builder.consumers {
		consume(employee)
	}
	return employee
}
