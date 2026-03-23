package main

import (
	"fmt"
	"slices"
)

func canPartition(nums []int) bool {
	// 先判断一下 target
	n := len(nums)
	sum := 0

	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	if sum%2 == 1 {
		return false
	}

	slices.Sort(nums)

	target := sum / 2

	// 构建一个dp数组，最终能让 dp[target]
	// 对于当前的元素有选和不选两种情况
	dp := make([]bool, target+1)

	for i := 0; i < n; i++ {
		dp[nums[i]] = true
		for j := target; j >= nums[i]; j-- {
			if dp[j-nums[i]] == true {
				dp[j] = true
			}
		}
	}

	return dp[target]
}

func main() {
	testCases := []struct {
		nums []int
		ans  bool
	}{
		{
			[]int{1, 5, 11, 5},
			true,
		},
		{
			[]int{1, 2, 3, 5},
			false,
		},
	}

	for _, testCase := range testCases {
		ans := canPartition(testCase.nums)
		fmt.Println(ans, testCase.ans == ans)
	}
}
