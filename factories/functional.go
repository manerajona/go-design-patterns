package factories

type Employee struct {
	name     string
	Position string
	Salary   int
}

func (e *Employee) Name() string {
	return e.name
}

func EmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}
