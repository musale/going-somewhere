package main

import "fmt"

const englishPrefix = "Hello, "
const spanish = "Spanish"
const spanishPrefix = "Hola, "
const french = "French"
const frenchPrefix = "Bonjour, "

func main() {
	fmt.Println(Hello("Kim", "English"))
}

// Hello returns a string
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishPrefix + name
	}
	if language == french {
		return frenchPrefix + name
	}
	return englishPrefix + name
}
