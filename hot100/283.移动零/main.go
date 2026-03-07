package main

import (
	"fmt"
	"reflect"
)

func moveZeroes(nums []int) {
	n := len(nums)
	left := 0

	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left += 1
		}
	}
}

func main() {
	testCases := []struct {
		nums []int
		ans  []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
	}

	for _, testCase := range testCases {
		moveZeroes(testCase.nums)
		// fmt.Print(testCase.ans, testCase.nums)
		fmt.Println(reflect.DeepEqual(testCase.nums, testCase.ans))
	}
}
