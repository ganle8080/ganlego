package main

import "testing"

func TestXxx(t *testing.T) {
	pivotIndex([]int{1, 7, 3, 6, 5, 6})
}

func pivotIndex(nums []int) int {
	resultInt := -1
	for i := 0; i < len(nums); i++ {
		leftResult := sumAll(nums[:i])
		rigthResult := sumAll(nums[i+1:])
		if leftResult == rigthResult {
			return i
		}
	}

	return resultInt
}

func sumAll(nums []int) int {
	var a = 0
	for _, num := range nums {
		a = a + num
	}

	return a
}
