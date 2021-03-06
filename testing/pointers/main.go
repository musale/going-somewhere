package main

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds is error thrown when you withdraw more than balance
var ErrInsufficientFunds = errors.New("cannot withdraw; balance is less than withdrawal amount")

func main() {}

// Stringer is a custom print formatter for our BTC
type Stringer struct {
	String string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

// Bitcoin is our currency
type Bitcoin float64

// Wallet stores our cash
type Wallet struct {
	balance Bitcoin
}

// Deposit amount into wallet
func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
	return
}

// Balance in the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw from the wallet
func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amt
	return nil
}
