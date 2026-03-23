package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	// dp[i] = dp[i - coins[j]] + 1
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

func main() {
	testCases := []struct {
		coins  []int
		amount int
		ans    int
	}{
		{
			[]int{1, 2, 5},
			11,
			3,
		},
		{
			[]int{2},
			3,
			-1,
		},
		{
			[]int{1},
			0,
			0,
		},
	}

	for _, testCase := range testCases {
		ans := coinChange(testCase.coins, testCase.amount)
		fmt.Println(ans, ans == testCase.ans)
	}
}
