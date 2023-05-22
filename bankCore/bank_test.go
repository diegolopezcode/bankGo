package bank

import (
	"testing"
)

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Diego",
			Address: "Cra 54 #23-12",
			Phone:   "83454345",
		},
		Number:  1011,
		Balance: 12243,
	}

	if account.Name == "" {
		t.Error("can't create an Account object")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Diego",
			Address: "Cra 54 #23-12",
			Phone:   "83454345",
		},
		Number:  1011,
		Balance: 0,
	}

	account.Deposit(10)

	if account.Balance != 10 {
		t.Error("Balance is not being updater after a deposit")
	}
}

func TestDepositInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Diego",
			Address: "Cra 54 #23-12",
			Phone:   "83454345",
		},
		Number:  1011,
		Balance: 0,
	}
	if err := account.Deposit(-5); err == nil {
		t.Error("only positive number should alowed to deposit")
	}

}

func TestWithDraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Diego",
			Address: "Cra 54 #23-12",
			Phone:   "83454345",
		},
		Number:  1011,
		Balance: 0,
	}

	account.Deposit(10)
	account.WithDraw(10)

	if account.Balance != 0 {
		t.Error("balance is not being updated after withdraw")
	}
}

func TestWithDrawInsufficient(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Diego",
			Address: "Cra 54 #23-12",
			Phone:   "83454345",
		},
		Number:  1011,
		Balance: 0,
	}

	account.Deposit(10)
	if err := account.WithDraw(11); err == nil {
		t.Error("the amount to withdraw should be greather than the account's balance")

	}

}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Diego",
			Address: "Cra 54 #23-12",
			Phone:   "83454345",
		},
		Number:  1011,
		Balance: 0,
	}

	account.Deposit(100)
	statement := account.Statement()
	if statement != "1011 - Diego - 100" {
		t.Error("statement doesn't have the proper format")
	}

}

//...

func TestTransfer(t *testing.T) {
	accountA := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	accountB := Account{
		Customer: Customer{
			Name:    "Mark",
			Address: "Irvine, California",
			Phone:   "(949) 555 0198",
		},
		Number:  1002,
		Balance: 0,
	}

	accountA.Deposit(100)
	err := accountA.Transfer(50, &accountB)

	if accountA.Balance != 50 && accountB.Balance != 50 {
		t.Error("transfer from account A to account B is not working", err)
	}
}
