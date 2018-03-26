package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrect := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("%s is not the same as %s", got, want)
		}
	}
	t.Run("print name passed in args", func(t *testing.T) {
		got := Hello("Thiel", "English")
		want := "Hello, Thiel"
		assertCorrect(t, got, want)
	})

	t.Run("print default when args is empty string", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrect(t, got, want)

	})

	t.Run("greet in spanish", func(t *testing.T) {
		got := Hello("Miley", "Spanish")
		want := "Hola, Miley"
		assertCorrect(t, got, want)
	})

	t.Run("greet in french", func(t *testing.T) {
		got := Hello("Miley", "French")
		want := "Bonjour, Miley"
		assertCorrect(t, got, want)
	})

}
