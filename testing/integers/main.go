package main

import "fmt"

func main() {
	sum := Add(2, 2)
	fmt.Println(sum)
}

// Add sums two integers
func Add(a, b int) int {
	return a + b
}
