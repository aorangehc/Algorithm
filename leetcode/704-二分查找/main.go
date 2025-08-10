package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func main() {
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9},
		{[]int{-1, 0, 3, 5, 9, 12}, 2},
		{[]int{-1, 0, 3, 5, 9, 12}, 13},
		{[]int{-12, 0, 3, 5, 9, 12}, -13},
	}

	for _, testCase := range testCases {
		fmt.Println(search(testCase.nums, testCase.target))
	}
}
