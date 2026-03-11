package main

import (
	"fmt"
	"reflect"
)

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}

	tmp := make([]int, n)

	copy(tmp, nums)

	for i := 0; i < n; i++ {
		nums[(i+k)%n] = tmp[i]
	}
}

func main() {
	testCases := []struct {
		nums []int
		k    int
		ans  []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			[]int{-1, -100, 3, 99},
			2,
			[]int{3, 99, -1, -100},
		},
	}

	for _, testCase := range testCases {
		rotate(testCase.nums, testCase.k)
		fmt.Println(testCase.nums, reflect.DeepEqual(testCase.ans, testCase.nums))
	}
}
