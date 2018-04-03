package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit into wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))

		got := wallet.Balance()
		want := Bitcoin(10.0)

		if got != want {
			t.Errorf("expected %s but got %s", want, got)
		}
	})

	t.Run("with from wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30.0)}

		wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(20.0)

		if got != want {
			t.Errorf("expected %s but got %s", want, got)
		}
	})
}
