package bank

import (
	"errors"
	"fmt"
)

type Customer struct {
	Name    string
	Address string
	Phone   string
}

type Account struct {
	Customer
	Balance float64
	Number  int32
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greather to Zero")
	}
	a.Balance += amount
	return nil
}

func (a *Account) WithDraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withDraw should be greather to Zero")
	}
	if amount > a.Balance {
		return errors.New("the amount to withdraw should be greather than the account's balance")
	}

	a.Balance -= amount
	return nil
}

// Statement
func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

// Transfer function
func (a *Account) Transfer(amount float64, dest *Account) error {
	if amount <= 0 {
		return errors.New("the amount to transfer should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to transfer should be greater than the account's balance")
	}

	a.WithDraw(amount)
	dest.Deposit(amount)
	return nil
}

// Bank ...
type Bank interface {
	Statement() string
}

// Statement ...
func Statement(b Bank) string {
	return b.Statement()
}
