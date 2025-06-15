package command

import (
	"errors"
	"testing"
)

func TestBankAccountCommand(t *testing.T) {
	type args struct {
		startBalance float64
		op           Operation
		amount       float64
	}
	tests := []struct {
		name            string
		args            args
		expectedErr     error
		expectedBalance float64
		revertBalance   float64
	}{
		{
			name:            "Deposit 100",
			args:            args{0, Deposit, 100},
			expectedErr:     nil,
			expectedBalance: 100,
			revertBalance:   0,
		},
		{
			name:            "Withdraw 50 from 100",
			args:            args{100, Withdraw, 50},
			expectedErr:     nil,
			expectedBalance: 50,
			revertBalance:   100,
		},
		{
			name:            "Withdraw over overdraft",
			args:            args{0, Withdraw, 1000},
			expectedErr:     errors.New("not enough funds"),
			expectedBalance: 0,
			revertBalance:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ba := NewBankAccount()
			_ = ba.Deposit(tt.args.startBalance)
			cmd := NewBankAccountCommand(ba, tt.args.op, tt.args.amount)

			succeeded := cmd.Handle()
			err := cmd.Error()
			if tt.expectedErr != nil {
				if succeeded {
					t.Errorf("Expected command to fail, but succeeded")
				}
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Errorf("Expected error %v, got %v", tt.expectedErr, err)
				}
				if ba.GetBalance() != tt.args.startBalance {
					t.Errorf("Expected balance %v, got %v", tt.args.startBalance, ba.GetBalance())
				}
				// Revert should be a no-op and not error
				reverted := cmd.Revert()
				if reverted {
					t.Errorf("Expected revert to do nothing and return false")
				}
				if ba.GetBalance() != tt.args.startBalance {
					t.Errorf("After revert, expected balance %v, got %v", tt.args.startBalance, ba.GetBalance())
				}
			} else {
				if !succeeded {
					t.Errorf("Expected command to succeed")
				}
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if ba.GetBalance() != tt.expectedBalance {
					t.Errorf("Expected balance %v, got %v", tt.expectedBalance, ba.GetBalance())
				}
				if !cmd.Succeeded() {
					t.Errorf("Expected command to be marked as succeeded")
				}
				// Test Revert
				reverted := cmd.Revert()
				if !reverted {
					t.Errorf("Expected revert to succeed and return true")
				}
				if ba.GetBalance() != tt.revertBalance {
					t.Errorf("After revert, expected balance %v, got %v", tt.revertBalance, ba.GetBalance())
				}
				if cmd.Succeeded() {
					t.Errorf("Expected command to be marked as not succeeded after revert")
				}
				// Retry Handle after revert should succeed again if possible
				retrySucceeded := cmd.Handle()
				if tt.args.op == Withdraw && tt.revertBalance < tt.args.amount {
					if retrySucceeded {
						t.Errorf("Retrying Handle should fail for insufficient funds")
					}
				} else {
					if !retrySucceeded {
						t.Errorf("Retrying Handle after revert should succeed if possible")
					}
				}
			}
		})
	}
}
