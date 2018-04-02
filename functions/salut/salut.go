package salut

import (
	"fmt"
)

// Salute structure
type Salute struct {
	Name     string
	Greeting string
}

// Greet functionality
func Greet(s Salute) {
	message, alternate := CreateMessages(s.Greeting, s.Name)
	fmt.Println(message)
	fmt.Printf(alternate)
}

// CreateMessages returns messages
func CreateMessages(greeting, name string) (message, alternate string) {
	message = greeting + " " + name
	alternate = "How are you " + name
	return
}
