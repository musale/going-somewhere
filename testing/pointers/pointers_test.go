package main

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("expected %s but got %s", want, got)
		}
	}

	t.Run("deposit into wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		want := Bitcoin(10.0)
		assertBalance(t, wallet, want)
	})

	t.Run("with from wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30.0)}
		wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(20.0)
		assertBalance(t, wallet, want)
	})
}
