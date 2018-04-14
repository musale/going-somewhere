package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Kendrick")

	got := buffer.String()
	want := "Hello, Kendrick"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
