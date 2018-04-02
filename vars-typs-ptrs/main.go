package main

import "fmt"

// SalutationHandler takes care of greetings
type SalutationHandler struct {
	name     string
	greeting string
}

func main() {
	// Types
	v1 := SalutationHandler{}
	v1.name = "King Kunta"
	v1.greeting = "Helloo"
	fmt.Println(v1.name, v1.greeting)

	v2 := SalutationHandler{name: "James", greeting: "James Bond"}
	fmt.Println(v2.name, v2.greeting)

	v3 := SalutationHandler{"Droid", "Hello King"}
	fmt.Println(v3.name, v3.greeting)

	// Variables
	message := "now we are talking"
	a, b, c := 123, false, "see"
	fmt.Println(message, a, b, c)

	// Pointers
	var messageVal = "Initial value"
	var messagePtr = &messageVal
	fmt.Println(messageVal, messagePtr)

	// Re-assign address value for message
	*messagePtr = "New message"
	fmt.Println(messageVal, messagePtr)
}
