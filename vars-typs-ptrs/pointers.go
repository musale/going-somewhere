package main

import (
	"fmt"
)

func main() {
	var message = "Initial value"
	var messagePtr = &message
	fmt.Println(message, messagePtr)

	// Re-assign address value for message
	*messagePtr = "New message"
	fmt.Println(message, messagePtr)
}
