package main

import "fmt"

func findXSum(nums []int, k, x int) []int64 {

}

func main() {
	testCases := []struct {
		nums []int
		k int
		x int
	} {
		{[]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2},
		{[]int{3, 8, 7, 8, 7, 5}, 2, 2},
	}

	for _, testCase := testCases {
		fmt.Println(findXSum(testCase.nums, testCase.k, testCase.x))
	}
}
