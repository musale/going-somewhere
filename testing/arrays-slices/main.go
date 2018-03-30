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

// SumAllTails returns sums of the last slice values except the first
func SumAllTails(numbersSlice ...[]int) []int {
	var sums []int
	for _, numbers := range numbersSlice {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
