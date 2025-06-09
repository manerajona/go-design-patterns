package builder

import (
	"testing"
)

func TestEmployeeBuilder_ValidInput(t *testing.T) {
	builder := EmployeeBuilder{}
	employee := builder.
		Name("Anne-Marie O'Neil").
		Phone("(123) 456-7890").
		Salary(45_100).
		Build()

	if employee.name != "Anne-Marie O'Neil" {
		t.Errorf("Expected name 'Anne-Marie O'Neil', got '%s'", employee.name)
	}
	if employee.phone != "(123) 456-7890" {
		t.Errorf("Expected phone '(123) 456-7890', got '%s'", employee.phone)
	}
	if employee.salary != 45100 {
		t.Errorf("Expected salary 45100, got %d", employee.salary)
	}
}

func TestEmployeeBuilder_InvalidName(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid name")
		}
	}()
	builder := EmployeeBuilder{}
	builder.Name("12345").Build()
}

func TestEmployeeBuilder_InvalidPhone(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid phone")
		}
	}()
	builder := EmployeeBuilder{}
	builder.Phone("12345").Build()
}

func TestEmployeeBuilder_InvalidSalaryLow(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for salary too low")
		}
	}()
	builder := EmployeeBuilder{}
	builder.Salary(10_000).Build()
}

func TestEmployeeBuilder_InvalidSalaryHigh(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for salary too high")
		}
	}()
	builder := EmployeeBuilder{}
	builder.Salary(2_000_000).Build()
}
