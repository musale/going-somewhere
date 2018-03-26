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
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}
