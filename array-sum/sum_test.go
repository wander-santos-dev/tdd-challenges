package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		wait := 15

		if got != wait {
			t.Errorf("Got %d, wait %d, given %v", got, wait, numbers)
		}
	})
}
