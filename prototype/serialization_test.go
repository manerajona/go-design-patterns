package prototype

import (
	"reflect"
	"testing"
)

func TestCompany_DeepCopy(t *testing.T) {
	original := &Company{
		Name:      "John & Johnson",
		HQ:        &HeadQuarters{"123 London Rd", "London", "UK"},
		Employees: []string{"Chris", "Matt", "Sam"},
	}
	dCopy := original.DeepCopy()

	// Ensure it's a deep copy
	if original == dCopy {
		t.Error("Expected a new pointer for company")
	}
	if original.HQ == dCopy.HQ {
		t.Error("Expected a new pointer for HQ")
	}
	if &original.Employees[0] == &dCopy.Employees[0] {
		t.Error("Expected a new slice for Employees")
	}
	if !reflect.DeepEqual(original, dCopy) {
		t.Errorf("Expected identical values, got\noriginal: %+v\ndCopy: %+v", original, dCopy)
	}

	// Mutate dCopy and ensure original is unaffected
	dCopy.Name = "HappyCorp"
	dCopy.HQ.StreetAddress = "321 Baker St"
	dCopy.Employees = append(dCopy.Employees, "Jill")

	if original.Name == dCopy.Name {
		t.Errorf("Original name changed when dCopy mutated")
	}
	if original.HQ.StreetAddress == dCopy.HQ.StreetAddress {
		t.Errorf("Original HQ changed when dCopy mutated")
	}
	if len(original.Employees) == len(dCopy.Employees) {
		t.Errorf("Original Employees changed when dCopy mutated")
	}
}
