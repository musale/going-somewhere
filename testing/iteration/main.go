package main

import "fmt"

func main() {
	rpt := Repeat("helloy", 5)
	fmt.Println(rpt)
}

// Repeat returns the word repeated 5 times
func Repeat(word string, times int) string {
	var rpt string
	for i := 0; i < times; i++ {
		rpt += word
	}
	return rpt
}
