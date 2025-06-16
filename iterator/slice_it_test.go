package iterator

import "testing"

func TestSliceIterator(t *testing.T) {
	data := []int{10, 20, 30}
	iter := NewSliceIterator(data)

	// Test HasNext, Next, and Current
	for i, expected := range data {
		if !iter.HasNext() {
			t.Fatalf("Expected HasNext to be true at index %d", i)
		}
		val, ok := iter.Next()
		if !ok {
			t.Fatalf("Expected Next to return true at index %d", i)
		}
		if val != expected {
			t.Errorf("Expected %d at index %d, got %d", expected, i, val)
		}
		if iter.Current() != expected {
			t.Errorf("Current() mismatch at index %d: got %d, want %d", i, iter.Current(), expected)
		}
	}

	// Test HasNext is false at end
	if iter.HasNext() {
		t.Errorf("Expected HasNext to be false at end")
	}

	// Test Next at end returns zero value and false
	val, ok := iter.Next()
	if ok {
		t.Errorf("Expected Next to return false at end")
	}
	if val != 0 {
		t.Errorf("Expected zero value at end, got %d", val)
	}

	// Test Reset
	iter.Reset()
	if iter.Current() != 0 {
		t.Errorf("Expected zero value after Reset, got %d", iter.Current())
	}
	val, ok = iter.Next()
	if !ok || val != 10 {
		t.Errorf("Expected Next after Reset to return 10 and true, got %d and %v", val, ok)
	}
}
