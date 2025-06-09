package solid

import (
	"reflect"
	"testing"
)

var (
	john  = Person{"John"}
	chris = Person{"Chris"}
	matt  = Person{"Matt"}
	sandy = Person{"Sandy"}
)

func TestRelationships_AddParentAndChildAndFindChildren(t *testing.T) {
	relationships := Relationships{}
	relationships.AddParentAndChild(&john, &chris)
	relationships.AddParentAndChild(&john, &matt)
	relationships.AddParentAndChild(&john, &sandy)

	children := relationships.FindChildrenOf(john)
	if len(children) != 3 {
		t.Fatalf("Expected 3 children, got %d", len(children))
	}
}

func TestRelationships_AddSiblingsAndFindSiblings(t *testing.T) {
	relationships := Relationships{}
	relationships.AddSiblings(&matt, &chris, &sandy)

	siblingsOfMatt := relationships.FindSiblingsOf(matt)
	if len(siblingsOfMatt) != 2 {
		t.Fatalf("Expected 2 siblings, got %d", len(siblingsOfMatt))
	}
}

func TestResearch_GenerateParentReport(t *testing.T) {
	relationships := Relationships{}
	relationships.AddParentAndChild(&john, &chris)
	research := Research{&relationships}

	report := research.GenerateParentReport(john)
	expected := []string{"John", "has a child called", "Chris"}
	if !reflect.DeepEqual(report, expected) {
		t.Errorf("Expected report %v, got %v", expected, report)
	}
}

func TestResearch_GenerateSiblingReport(t *testing.T) {
	relationships := Relationships{}
	relationships.AddSiblings(&matt, &chris, &sandy)
	research := Research{&relationships}

	report := research.GenerateSiblingReport(matt)
	// Siblings: Chris and Sandy
	expected := []string{"Matt", "has a sibling called", "Chris", "Matt", "has a sibling called", "Sandy"}
	if !reflect.DeepEqual(report, expected) {
		t.Errorf("Expected report %v, got %v", expected, report)
	}
}
