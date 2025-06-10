package prototype

import (
	"testing"
)

func TestEmployee_DeepCopy(t *testing.T) {
	emp := &Employee{"Alice", Location{101, "123 East Dr", "London"}}
	empCopy := emp.DeepCopy()
	if emp == empCopy {
		t.Error("Expected DeepCopy to return a different pointer")
	}
	if emp.Name != empCopy.Name || emp.Office != empCopy.Office {
		t.Errorf("Expected identical field values: got %+v vs %+v", emp, empCopy)
	}
	empCopy.Name = "Bob"
	empCopy.Office.Suite = 202
	if emp.Name == empCopy.Name {
		t.Errorf("Name modified in original after DeepCopy")
	}
	if emp.Office.Suite == empCopy.Office.Suite {
		t.Errorf("Office.Suite modified in original after DeepCopy")
	}
}

func TestNewLondonOfficeEmployee(t *testing.T) {
	john := NewLondonOfficeEmployee("John", 100)
	if john.Name != "John" {
		t.Errorf("Expected name 'John', got %q", john.Name)
	}
	if john.Office.StreetAddress != "123 East Dr" || john.Office.City != "London" {
		t.Errorf("Expected London office, got %+v", john.Office)
	}
	if john.Office.Suite != 100 {
		t.Errorf("Expected suite 100, got %d", john.Office.Suite)
	}
}

func TestNewParisOfficeEmployee(t *testing.T) {
	jane := NewParisOfficeEmployee("Jane", 200)
	if jane.Name != "Jane" {
		t.Errorf("Expected name 'Jane', got %q", jane.Name)
	}
	if jane.Office.StreetAddress != "66 Victor Hugo Ave" || jane.Office.City != "Paris" {
		t.Errorf("Expected Paris office, got %+v", jane.Office)
	}
	if jane.Office.Suite != 200 {
		t.Errorf("Expected suite 200, got %d", jane.Office.Suite)
	}
}
