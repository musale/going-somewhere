package main

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	sum := SumAll([]int{1, 2, 3}, []int{4, 5})
	expected := []int{6, 9}
	if !reflect.DeepEqual(sum, expected) {
		t.Errorf("expected %v but got %v", expected, sum)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, expected, sum []int) {
		if !reflect.DeepEqual(sum, expected) {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	}
	t.Run("sum tails of a slice", func(t *testing.T) {
		sum := SumAllTails([]int{1, 2, 3}, []int{4, 5})
		expected := []int{5, 5}
		checkSums(t, expected, sum)
	})

	t.Run("sum tails with empty slice", func(t *testing.T) {
		sum := SumAllTails([]int{}, []int{4, 5})
		expected := []int{0, 5}
		checkSums(t, expected, sum)
	})
}
