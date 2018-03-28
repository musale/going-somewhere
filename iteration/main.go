package main

import "fmt"

func main() {
	rpt := Repeat("helloy")
	fmt.Println(rpt)
}

// Repeat returns the word repeated 5 times
func Repeat(word string) string {
	var rpt string
	for i := 0; i < 5; i++ {
		rpt += word
	}
	return rpt
}
