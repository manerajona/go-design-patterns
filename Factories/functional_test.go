package Factories

import (
	"testing"
)

func TestEmployeeFactory(t *testing.T) {
	devFactory := EmployeeFactory("Developer", 60000)

	adam := devFactory("Adam")
	if adam.Name() != "Adam" {
		t.Errorf("Expected name 'Adam', got %q", adam.Name())
	}
	if adam.Position != "Developer" {
		t.Errorf("Expected position 'Developer', got %q", adam.Position)
	}
	if adam.Salary != 60000 {
		t.Errorf("Expected salary 60000, got %d", adam.Salary)
	}
}

func TestEmployeeFactory_UpdateSalary(t *testing.T) {
	devFactory := EmployeeFactory("Developer", 60000)
	adam := devFactory("Adam")
	adam.Salary += 10000
	if adam.Salary != 70000 {
		t.Errorf("Expected updated salary 70000, got %d", adam.Salary)
	}
}

func TestEmployeeFactory_UpdatePosition(t *testing.T) {
	managerFactory := EmployeeFactory("Manager", 80000)
	jane := managerFactory("Jane")
	jane.Position = "Vice President"
	if jane.Position != "Vice President" {
		t.Errorf("Expected position 'Vice President', got %q", jane.Position)
	}
}
