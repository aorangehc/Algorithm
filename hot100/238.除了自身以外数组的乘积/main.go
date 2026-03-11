package main

import (
	"fmt"
	"reflect"
)

func productExceptSelf(nums []int) []int {
	// 两次遍历，构建上三角和下三角
	n := len(nums)
	ans := make([]int, n)
	up := make([]int, n)
	down := make([]int, n)
	up[0] = 1

	for i := 1; i < n; i++ {
		up[i] = up[i-1] * nums[i-1]
	}

	down[n-1] = 1

	for i := n - 2; i >= 0; i-- {
		down[i] = down[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		ans[i] = up[i] * down[i]
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		ans  []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{24, 12, 8, 6},
		},
		{
			[]int{-1, 1, 0, -3, 3},
			[]int{0, 0, 9, 0, 0},
		},
	}

	for _, testCase := range testCases {
		ans := productExceptSelf(testCase.nums)
		fmt.Println(ans, reflect.DeepEqual(ans, testCase.ans))
	}
}
