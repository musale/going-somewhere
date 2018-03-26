package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrect := func(t *testing.T, got, want string) {
		if got != want {
			t.Errorf("%s is not the same as %s", got, want)
		}
	}
	t.Run("print name passed in args", func(t *testing.T) {
		got := Hello("Thiel")
		want := "Hello, Thiel"
		assertCorrect(t, got, want)
	})

	t.Run("print default when args is empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrect(t, got, want)

	})

}
