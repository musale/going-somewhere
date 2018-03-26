package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Thiel")
	want := "Hello, Thiel"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
