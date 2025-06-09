package builder

import (
	"testing"
)

func TestPersonBuilder_FluentAPI(t *testing.T) {
	person :=
		NewPersonBuilder().
			Name("John").
			Lives().At("100 Almond St.").In("London").WithPostcode("SW12BC").
			Works().At("CodeCo").AsA("Engineer").Earning(100_000).
			Build()

	if person.Name != "John" {
		t.Errorf("Expected Name 'John', got '%s'", person.Name)
	}
	if person.StreetAddress != "100 Almond St." {
		t.Errorf("Expected StreetAddress '100 Almond St.', got '%s'", person.StreetAddress)
	}
	if person.City != "London" {
		t.Errorf("Expected City 'London', got '%s'", person.City)
	}
	if person.Postcode != "SW12BC" {
		t.Errorf("Expected Postcode 'SW12BC', got '%s'", person.Postcode)
	}
	if person.CompanyName != "CodeCo" {
		t.Errorf("Expected CompanyName 'CodeCo', got '%s'", person.CompanyName)
	}
	if person.Position != "Engineer" {
		t.Errorf("Expected Position 'Engineer', got '%s'", person.Position)
	}
	if person.AnnualIncome != 100_000 {
		t.Errorf("Expected AnnualIncome 100000, got %d", person.AnnualIncome)
	}
}

func TestPersonBuilder_OnlyAddress(t *testing.T) {
	person :=
		NewPersonBuilder().
			Name("Alice").
			Lives().At("1 Main St.").In("Paris").WithPostcode("75001").
			Build()

	if person.Name != "Alice" || person.StreetAddress != "1 Main St." ||
		person.City != "Paris" || person.Postcode != "75001" {
		t.Errorf("Person address building failed: %+v", *person)
	}
}

func TestPersonBuilder_OnlyJob(t *testing.T) {
	person :=
		NewPersonBuilder().
			Name("Bob").
			Works().At("MegaCorp").AsA("Manager").Earning(200_000).
			Build()

	if person.Name != "Bob" || person.CompanyName != "MegaCorp" ||
		person.Position != "Manager" || person.AnnualIncome != 200_000 {
		t.Errorf("Person job building failed: %+v", *person)
	}
}
