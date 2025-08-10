package main

import "fmt"

func maxSubArray(nums []int) int {
	ans := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i]) // 以第i个元素结尾的最大连续子数组的和，这里必须包含第i个元素
		ans = max(ans, dp[i])
	}

	return ans
}

func main() {
	testCases := [][]int{
		[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		[]int{1},
		[]int{5, 4, -1, 7, 8},
	}

	for _, testCase := range testCases {
		ans := maxSubArray(testCase)
		fmt.Println(ans)
	}
}
