package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit into wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		want := Bitcoin(10.0)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw from wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30.0)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(20.0)
		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("withdraw more than deposit", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(50))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()

	if got != want {
		t.Errorf("expected %s but got %s", want, got)
	}
}

func assertError(t *testing.T, err error, want error) {
	if err == nil {
		t.Fatal("wanted an error but didnt get one")
	}

	if err != want {
		t.Errorf("expected %s but got %s", want, err)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}
}
