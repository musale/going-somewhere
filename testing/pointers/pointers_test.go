package main

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("expected %s but got %s", want, got)
		}
	}
	assertError := func(t *testing.T, err error, want string) {
		if err == nil {
			t.Fatal("wanted an error but didnt get one")
		}

		if err.Error() != want {
			t.Errorf("expected %s but got %s", want, err.Error())
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

	t.Run("withdraw more than deposit", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(50))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "cannot withdraw; balance is less than withdrawal amount")
	})
}
