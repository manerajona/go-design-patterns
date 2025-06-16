package memento

import "testing"

func TestBankAccount_PushPullUndoRedo(t *testing.T) {
	ba := NewBankAccount(100)

	if got := ba.GetBalance(); got != 100 {
		t.Errorf("Expected initial balance 100, got %v", got)
	}

	ba.Push(50) // 150
	if got := ba.GetBalance(); got != 150 {
		t.Errorf("Expected balance 150 after deposit, got %v", got)
	}

	ba.Push(25) // 175
	if got := ba.GetBalance(); got != 175 {
		t.Errorf("Expected balance 175 after deposit, got %v", got)
	}

	ba.Pull(10) // 165
	if got := ba.GetBalance(); got != 165 {
		t.Errorf("Expected balance 165 after withdraw, got %v", got)
	}

	// Undo Withdraw (back to 175)
	ba.Undo()
	if got := ba.GetBalance(); got != 175 {
		t.Errorf("Expected balance 175 after undo, got %v", got)
	}
	// Undo Deposit (back to 150)
	ba.Undo()
	if got := ba.GetBalance(); got != 150 {
		t.Errorf("Expected balance 150 after undo, got %v", got)
	}
	// Undo Deposit (back to 100)
	ba.Undo()
	if got := ba.GetBalance(); got != 100 {
		t.Errorf("Expected balance 100 after undo, got %v", got)
	}
	// Undo at initial state should not go below 0th offset
	ba.Undo()
	if got := ba.GetBalance(); got != 100 {
		t.Errorf("Expected balance 100 after extra undo, got %v", got)
	}

	// Redo Deposit (back to 150)
	ba.Redo()
	if got := ba.GetBalance(); got != 150 {
		t.Errorf("Expected balance 150 after redo, got %v", got)
	}
	// Redo Deposit (back to 175)
	ba.Redo()
	if got := ba.GetBalance(); got != 175 {
		t.Errorf("Expected balance 175 after redo, got %v", got)
	}
	// Redo Withdraw (back to 165)
	ba.Redo()
	if got := ba.GetBalance(); got != 165 {
		t.Errorf("Expected balance 165 after redo, got %v", got)
	}
	// Redo at latest state should not go past the end
	ba.Redo()
	if got := ba.GetBalance(); got != 165 {
		t.Errorf("Expected balance 165 after extra redo, got %v", got)
	}
}
