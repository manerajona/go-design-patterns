package factories

import (
	"testing"
)

func TestNewPerson_Valid(t *testing.T) {
	p, err := NewPerson("Jane", 21)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if p.Name() != "Jane" {
		t.Errorf("Expected name 'Jane', got %q", p.Name())
	}
	if p.Age() != 21 {
		t.Errorf("Expected age 21, got %d", p.Age())
	}
}

func TestNewPerson_BlankName(t *testing.T) {
	_, err := NewPerson("", 20)
	if err == nil {
		t.Fatalf("Expected error for blank name, got nil")
	}
}

func TestNewPerson_NegativeAge(t *testing.T) {
	_, err := NewPerson("John", -1)
	if err == nil {
		t.Fatalf("Expected error for negative age, got nil")
	}
}
