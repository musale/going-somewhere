package main

import (
	"bytes"
	"fmt"
)

func main() {}

// Greet prints name passed.
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
