package state

import (
	"testing"
)

func TestLoginSuccess(t *testing.T) {
	acc := NewAccount("bob", "bobpass")
	err := acc.Login("bobpass")
	if err != nil {
		t.Errorf("expected successful login, got error: %v", err)
	}
	if acc.State != Unlocked {
		t.Errorf("expected state Unlocked after login, got %v", acc.State)
	}
	if acc.attempts != 0 {
		t.Errorf("expected 0 attempts after success, got %d", acc.attempts)
	}
}

func TestLoginIncorrectPassword(t *testing.T) {
	acc := NewAccount("carol", "carolpass")
	for i := 1; i <= 2; i++ {
		err := acc.Login("wrongpass")
		if err == nil || err.Error() != "incorrect password" {
			t.Errorf("expected 'incorrect password', got %v", err)
		}
		if acc.State != Locked {
			t.Errorf("expected state Locked, got %v", acc.State)
		}
		if acc.attempts != i {
			t.Errorf("expected %d attempts, got %d", i, acc.attempts)
		}
	}
}

func TestLoginBlockedAfterThreeFails(t *testing.T) {
	acc := NewAccount("dan", "danpass")
	for i := 1; i <= 2; i++ {
		acc.Login("wrong")
	}
	err := acc.Login("wrong")
	if err == nil || err.Error() != "your account has been blocked, contact support" {
		t.Errorf("expected blocked account error, got %v", err)
	}
	if acc.State != Blocked {
		t.Errorf("expected state Blocked, got %v", acc.State)
	}
}

func TestBlockedAccountCannotLogin(t *testing.T) {
	acc := NewAccount("eve", "evepass")
	// block the account
	for i := 0; i < 3; i++ {
		acc.Login("wrong")
	}
	err := acc.Login("evepass")
	if err == nil || err.Error() != "account is blocked, contact support" {
		t.Errorf("expected blocked account error, got %v", err)
	}
	if acc.State != Blocked {
		t.Errorf("expected state Blocked, got %v", acc.State)
	}
}

func TestLoginSuccessResetsAttempts(t *testing.T) {
	acc := NewAccount("frank", "frankpass")
	acc.Login("wrong")
	acc.Login("frankpass")
	if acc.attempts != 0 {
		t.Errorf("expected attempts reset to 0 after success, got %d", acc.attempts)
	}
}

func TestUnlockedAccountLoginIsNoop(t *testing.T) {
	acc := NewAccount("gina", "ginapass")
	acc.Login("ginapass")
	err := acc.Login("ginapass")
	if err != nil {
		t.Errorf("expected nil error on login when already unlocked, got %v", err)
	}
}

func TestLogoutResetsStateToLocked(t *testing.T) {
	acc := NewAccount("hugo", "hugopass")
	acc.Login("hugopass")
	acc.Logout()
	if acc.State != Locked {
		t.Errorf("expected state Locked after logout, got %v", acc.State)
	}
}
