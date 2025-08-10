package main

import "fmt"

func moveZeroes(nums []int) {
	left := 0

	for i := range len(nums) {
		if nums[i] != 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left += 1
		}
	}
}

func main() {
	testCases := [][]int{
		[]int{0, 1, 0, 3, 12},
		[]int{0},
	}

	for _, testCase := range testCases {
		moveZeroes(testCase)
		fmt.Println(testCase)
	}
}
