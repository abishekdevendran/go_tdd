package arrays

import (
	"reflect"
	"testing"
)

func Sum(arr []int) int {
	sum := 0
	for _, el := range arr {
		sum += el
	}
	return sum
}

func SumAll(arrs ...[]int) []int {
	sums := make([]int, len(arrs))
	for i, el := range arrs {
		sums[i] = Sum(el)
	}
	return sums
}

func TestSum(t *testing.T) {

	t.Run("Individual Sum", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Slices sum", func(t *testing.T) {

		got := SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4 , 5})
		want := []int{6, 15}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
