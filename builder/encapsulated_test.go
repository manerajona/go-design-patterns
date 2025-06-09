package builder

import (
	"fmt"
	"testing"
)

type DummyVendor struct {
}

func (vendor *DummyVendor) Send(e *email) error {
	fmt.Println("Sending email", *e)
	return nil
}

func TestEmailBuilder_FluentAPI(t *testing.T) {
	builder := &EmailBuilder{}
	builder.
		From("alice@example.com").
		To("bob@example.com").
		Subject("Hello").
		Body("How are you?")

	if builder.from.Address != "alice@example.com" {
		t.Errorf("From: want %q, got %q", "alice@example.com", builder.from.Address)
	}
	if builder.to.Address != "bob@example.com" {
		t.Errorf("To: want %q, got %q", "bob@example.com", builder.to.Address)
	}
	if builder.subject != "Hello" {
		t.Errorf("Subject: want %q, got %q", "Hello", builder.subject)
	}
	if builder.body != "How are you?" {
		t.Errorf("Body: want %q, got %q", "How are you?", builder.body)
	}
}

func TestEmailBuilder_EmptySubjectPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for empty subject, got none")
		}
	}()
	builder := &EmailBuilder{}
	builder.Subject("")
}

func TestEmailBuilder_EmptyBodyPanics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for empty body, got none")
		}
	}()
	builder := &EmailBuilder{}
	builder.Body("")
}

func TestEmailBuilder_InvalidFromEmail(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid email")
		}
	}()
	builder := &EmailBuilder{}
	builder.From("example.com")
}

func TestEmailBuilder_InvalidToEmail(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for invalid email")
		}
	}()
	builder := &EmailBuilder{}
	builder.To("example.com")
}

func TestSendEmail_CallsVendor(t *testing.T) {
	vendor := &DummyVendor{}
	err := SendEmail(func(builder *EmailBuilder) {
		builder.
			From("foo@bar.com").
			To("bar@baz.com").
			Subject("Meeting").
			Body("Hello, do you want to meet?")
	}, vendor)

	if err != nil {
		t.Fatalf("SendEmail returned unexpected error: %v", err)
	}
}
