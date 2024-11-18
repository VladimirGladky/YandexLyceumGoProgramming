package main

import (
	"errors"
	"fmt"
)

type Account struct {
	balance float64
	owner   string
}

func NewAccount(owner string) *Account {
	return &Account{0.0, owner}
}

func (a *Account) SetBalance(bal float64) error {
	if bal >= 0.0 {
		a.balance = bal
		return nil
	}
	return errors.New("отрицательный баланс")
}

func (a *Account) Deposit(money float64) error {
	if a.balance+money >= 0 {
		a.balance = a.balance + money
		return nil
	}
	return errors.New("отрицательный баланс")
}

func (a *Account) Withdraw(money float64) error {
	if a.balance-money >= 0 && money >= 0 {
		a.balance = a.balance - money
		return nil
	}
	return errors.New("отрицательный баланс")
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func main() {
	account := NewAccount("Alice")
	fmt.Println(account.owner)
}
