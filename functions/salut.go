package main

import (
	"fmt"
)

// Salute structure
type Salute struct {
	name     string
	greeting string
}

// Greet functionality
func Greet(s Salute) {
	message, alternate := CreateMessages(s.greeting, s.name)
	fmt.Println(message)
	fmt.Printf(alternate)
}

// CreateMessages returns messages
func CreateMessages(greeting, name string) (message, alternate string) {
	message = greeting + " " + name
	alternate = "How are you " + name
	return
}

func main() {
	s := Salute{name: "Mr King", greeting: "Ni hao"}
	Greet(s)
}
