package state

import (
	"errors"
)

type State int

const (
	Locked State = iota
	Unlocked
	Blocked
)

type User struct {
	Username, Password string
}

type Account struct {
	User
	State
	attempts int
}

func NewAccount(username, password string) *Account {
	return &Account{
		User:  User{username, password},
		State: Locked,
	}
}

func (a *Account) Login(password string) error {
	switch a.State {
	case Unlocked:
		return nil
	case Locked:
		if password == a.Password {
			a.State = Unlocked
			a.attempts = 0 // reset attempts on success
			return nil
		}
		a.attempts++
		if a.attempts >= 3 {
			a.State = Blocked
			return errors.New("your account has been blocked, contact support")
		}
		return errors.New("incorrect password")
	case Blocked:
		return errors.New("account is blocked, contact support")
	}
	return errors.New("invalid state")
}

func (a *Account) Logout() {
	a.State = Locked
}
