package prototype

import (
	"reflect"
	"testing"
)

func TestAddress_DeepCopy(t *testing.T) {
	addr := &Address{"123 London Rd", "London", "UK"}
	addrCopy := addr.DeepCopy()
	if addr == addrCopy {
		t.Error("Expected a different pointer for copied address")
	}
	if *addr != *addrCopy {
		t.Errorf("Expected identical contents: %v vs %v", *addr, *addrCopy)
	}
	addrCopy.StreetAddress = "New Street"
	if addr.StreetAddress == addrCopy.StreetAddress {
		t.Errorf("Modification to copy affected original")
	}
}

func TestPerson_DeepCopy(t *testing.T) {
	john := &Person{
		Name:    "John",
		Address: &Address{"123 London Rd", "London", "UK"},
		Friends: []string{"Chris", "Matt"},
	}
	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	// Check that original is not affected
	if john.Name == jane.Name {
		t.Errorf("Original name modified")
	}
	if john.Address.StreetAddress == jane.Address.StreetAddress {
		t.Errorf("Original address modified")
	}
	if reflect.DeepEqual(john.Friends, jane.Friends) {
		t.Errorf("Original friends modified")
	}
	// Check original remains as expected
	if john.Name != "John" {
		t.Errorf("Expected John, got %s", john.Name)
	}
	if john.Address.StreetAddress != "123 London Rd" {
		t.Errorf("Address changed in original")
	}
	if len(john.Friends) != 2 || john.Friends[0] != "Chris" || john.Friends[1] != "Matt" {
		t.Errorf("Friends changed in original: %v", john.Friends)
	}
}
