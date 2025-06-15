package command

import (
	"errors"
)

var overdraftLimit = -500

type BankAccount interface {
	GetBalance() float64
	Deposit(float64) error
	Withdraw(float64) error
}

type bankAccount struct {
	balance float64
}

func (b *bankAccount) GetBalance() float64 {
	return b.balance
}

func (b *bankAccount) Deposit(amount float64) error {
	b.balance += amount
	return nil
}

func (b *bankAccount) Withdraw(amount float64) error {
	if b.balance-amount >= float64(overdraftLimit) {
		b.balance -= amount
		return nil
	}
	return errors.New("not enough funds")
}

func NewBankAccount() BankAccount {
	return &bankAccount{}
}

type Command interface {
	Handle() bool
	Revert() bool
	Succeeded() bool
	Error() error
}

type Operation int

const (
	Deposit Operation = iota
	Withdraw
)

type BankAccountCommand struct {
	account   BankAccount
	operation Operation
	amount    float64
	succeeded bool
	err       error
}

func (c *BankAccountCommand) Handle() bool {
	// Only execute if not succeeded
	if !c.succeeded {
		switch c.operation {
		case Deposit:
			c.err = c.account.Deposit(c.amount)
		case Withdraw:
			c.err = c.account.Withdraw(c.amount)
		default:
			c.err = errors.New("invalid operation")
		}
		c.succeeded = c.err == nil
	}
	return c.succeeded
}

func (c *BankAccountCommand) Revert() bool {
	// Only revert if previously succeeded
	if c.succeeded {
		var revertErr error
		switch c.operation {
		case Deposit:
			revertErr = c.account.Withdraw(c.amount)
		case Withdraw:
			revertErr = c.account.Deposit(c.amount)
		default:
			revertErr = errors.New("invalid operation")
		}
		// Only update state if revert was successful
		if revertErr == nil {
			c.succeeded = false
			c.err = nil // allow retrying Handle after revert
			return true
		}
	}
	return false
}

func (c *BankAccountCommand) Succeeded() bool {
	return c.succeeded
}

func (c *BankAccountCommand) Error() error {
	return c.err
}

func NewBankAccountCommand(account BankAccount, action Operation, amount float64) Command {
	return &BankAccountCommand{account: account, operation: action, amount: amount}
}
