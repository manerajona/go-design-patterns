package chain_of_resp

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func TestCommandModifier(t *testing.T) {
	type args struct {
		actionCalled bool
		role         Role
	}
	tests := []struct {
		name                 string
		args                 args
		expectedActionCalled bool
		expectedError        error
	}{
		{
			name: "Contributor role - should succeed",
			args: args{
				actionCalled: false,
				role:         Contributor,
			},
			expectedActionCalled: true,
			expectedError:        nil,
		},
		{
			name: "Admin role - should succeed",
			args: args{
				actionCalled: false,
				role:         Admin,
			},
			expectedActionCalled: true,
			expectedError:        nil,
		},
		{
			name: "Viewer role - should fail authorization",
			args: args{
				actionCalled: false,
				role:         Viewer,
			},
			expectedActionCalled: false,
			expectedError:        errors.New("invalid role 0"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actionCalled := tt.args.actionCalled
			cmd := NewCommand("TestUser", tt.args.role, func() error { actionCalled = true; return nil })

			root := NewCommandModifier(cmd)
			root.Add(NewAuthorizationCommandModifier(cmd))
			root.Add(NewCorrelationIdCommandModifier(cmd))
			root.Add(NewTimestampCommandModifier(cmd))

			err := root.Handle()

			if tt.expectedError != nil {
				if err == nil {
					t.Errorf("Expected error %v, got nil", tt.expectedError)
				} else if !strings.Contains(err.Error(), tt.expectedError.Error()) {
					t.Errorf("Expected error to contain %q, got %q", tt.expectedError.Error(), err.Error())
				}
				if actionCalled {
					t.Errorf("Expected action NOT to be called, but it was")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, got %v", err)
				}
				if !actionCalled {
					t.Errorf("Expected action to be called, but it was not")
				}
				if cmd.CorrelationId == "" {
					t.Error("Expected CorrelationId to be set")
				}
				if time.Since(cmd.Timestamp) > 10*time.Second {
					t.Errorf("Expected recent Timestamp, got %v", cmd.Timestamp)
				}
			}
		})
	}
}
