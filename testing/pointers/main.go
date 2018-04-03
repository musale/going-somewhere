package main

func main() {}

// Wallet stores our cash
type Wallet struct {
	balance float64
}

// Deposit amount into wallet
func (w Wallet) Deposit(amt float64) {
	w.balance += amt
	return
}

// Balance in the wallet
func (w Wallet) Balance() float64 {
	return w.balance
}
