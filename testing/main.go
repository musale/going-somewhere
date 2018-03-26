package main

import "fmt"

func main() {
	fmt.Println(Hello("Kim"))
}

// Hello returns a string
func Hello(name string) string {
	return "Hello " + name
}
