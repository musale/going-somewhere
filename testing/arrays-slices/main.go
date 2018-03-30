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
	lengthOfNumbers := len(numbers)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbers {
		sums[i] = Sum(numbers)
	}

	return sums
}
