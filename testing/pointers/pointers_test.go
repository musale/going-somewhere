package main

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10.0))

	got := wallet.Balance()
	want := Bitcoin(10.0)

	if got != want {
		t.Errorf("expected %s but got %s", want, got)
	}
}
