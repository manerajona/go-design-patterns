package chain_of_resp

import (
	"fmt"
	"time"
)

type Role int

const (
	Viewer Role = iota
	Contributor
	Admin
)

type Action func() error

type Command struct {
	From string
	Role
	CorrelationId string
	Timestamp     time.Time
	Request       Action
}

func NewCommand(from string, role Role, request Action) *Command {
	return &Command{From: from, Role: role, Request: request}
}

type Modifier interface {
	Add(m Modifier)
	Handle() error
}

type CommandModifier struct {
	command *Command
	next    Modifier // single linked list
}

func (cm *CommandModifier) Add(m Modifier) {
	if cm.next != nil {
		cm.next.Add(m)
	} else {
		cm.next = m
	}
}

func (cm *CommandModifier) Handle() error {
	if cm.next != nil {
		return cm.next.Handle()
	}
	// Handle the request
	return cm.command.Request()
}

func NewCommandModifier(command *Command) *CommandModifier {
	return &CommandModifier{command: command}
}

type CorrelationIdCommandModifier struct {
	CommandModifier
}

func NewCorrelationIdCommandModifier(c *Command) *CorrelationIdCommandModifier {
	return &CorrelationIdCommandModifier{CommandModifier{command: c}}
}

func (cm *CorrelationIdCommandModifier) Handle() error {
	cm.command.CorrelationId = fmt.Sprintf("%d", time.Now().UnixNano())
	return cm.CommandModifier.Handle()
}

type TimestampCommandModifier struct {
	CommandModifier
}

func NewTimestampCommandModifier(c *Command) *TimestampCommandModifier {
	return &TimestampCommandModifier{CommandModifier{command: c}}
}

func (cm *TimestampCommandModifier) Handle() error {
	cm.command.Timestamp = time.Now()
	return cm.CommandModifier.Handle()
}

type AuthorizationCommandModifier struct {
	CommandModifier
}

func NewAuthorizationCommandModifier(c *Command) *AuthorizationCommandModifier {
	return &AuthorizationCommandModifier{CommandModifier{command: c}}
}

func (cm *AuthorizationCommandModifier) Handle() error {
	if cm.command.Role < Contributor {
		return fmt.Errorf("invalid role %d", cm.command.Role)
	}
	return cm.CommandModifier.Handle()
}
