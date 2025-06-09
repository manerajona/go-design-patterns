package builder

import (
	"net/mail"
)

type email struct {
	from, to      mail.Address
	subject, body string
}

type EmailBuilder struct {
	email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	address, err := mail.ParseAddress(from)
	if err != nil {
		panic(err)
	}
	b.from = *address
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	address, err := mail.ParseAddress(to)
	if err != nil {
		panic(err)
	}
	b.to = *address
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	if subject == "" {
		panic("Email subject can't be empty")
	}
	b.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	if body == "" {
		panic("Body can't be empty")
	}
	b.email.body = body
	return b
}

type EmailVendor interface {
	Send(e *email) error
}

type EmailBuild func(*EmailBuilder)

func SendEmail(build EmailBuild, vendor EmailVendor) error {
	builder := EmailBuilder{}
	build(&builder)
	return vendor.Send(&builder.email)
}
