package main

import "fmt"

func main() {
	sum := Sum([5]int{1, 2, 3, 4, 5})
	fmt.Println(sum)
}

// Sum adds numbers in an array
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
