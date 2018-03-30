package main

func main() {
	Sum([]int{1, 2, 3, 4, 5})
}

// Sum adds numbers in an array
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll returns a slice of sums
func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, numbers := range numbers {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
