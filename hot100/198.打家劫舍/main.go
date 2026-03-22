package main

import "fmt"

func rob(nums []int) int {
	// 选或不选的问题
	// 选当前的话，上一个就不能选
	// dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
	n := len(nums)

	dp := make([]int, n)

	dp[0] = nums[0]

	if n >= 2 {
		dp[1] = max(dp[0], nums[1])
	}

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}

func main() {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			[]int{1, 2, 3, 1},
			4,
		},
		{
			[]int{2, 7, 9, 3, 1},
			12,
		},
	}

	for _, testCase := range testCases {
		ans := rob(testCase.nums)
		fmt.Println(ans, testCase.ans == ans)
	}
}
