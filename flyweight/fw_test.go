package flyweight

import (
	"strings"
	"testing"
)

func tearDown() {
	nameRepo = nil
}

func TestNewUser_FullName(t *testing.T) {
	// Reset global state when finished
	defer tearDown()

	u1 := NewUser("John Doe")
	u2 := NewUser("Jane Doe")
	u3 := NewUser("Jane Smith")

	if u1.FullName() != "John Doe" {
		t.Errorf("Expected 'John Doe', got '%s'", u1.FullName())
	}
	if u2.FullName() != "Jane Doe" {
		t.Errorf("Expected 'Jane Doe', got '%s'", u2.FullName())
	}
	if u3.FullName() != "Jane Smith" {
		t.Errorf("Expected 'Jane Smith', got '%s'", u3.FullName())
	}
	if len(nameRepo) != 4 {
		t.Errorf("Expected 4 unique names in nameRepo, got %d: %v", len(nameRepo), nameRepo)
	}
}

func TestNewUser_SharedNames(t *testing.T) {
	// Reset global state when finished
	defer tearDown()

	_ = NewUser("Alice Bob")
	_ = NewUser("Alice Charlie")
	_ = NewUser("Alice Bob")
	_ = NewUser("Charlie Bob")
	// nameRepo should have "Alice", "Bob", "Charlie"
	if len(nameRepo) != 3 {
		t.Errorf("Expected 3 unique names in nameRepo, got %d: %v", len(nameRepo), nameRepo)
	}
	if !strings.Contains(strings.Join(nameRepo, ","), "Alice") {
		t.Error("Expected 'Alice' in nameRepo")
	}
	if !strings.Contains(strings.Join(nameRepo, ","), "Bob") {
		t.Error("Expected 'Bob' in nameRepo")
	}
	if !strings.Contains(strings.Join(nameRepo, ","), "Charlie") {
		t.Error("Expected 'Charlie' in nameRepo")
	}
}
