package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any sized ints", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		expected := 6
		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	})
}
