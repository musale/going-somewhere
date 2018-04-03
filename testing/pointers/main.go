package main

func main() {}

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
