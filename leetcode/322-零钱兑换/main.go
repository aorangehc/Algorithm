package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // 表示凑成i需要的最少零钱数量

	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin] < amount {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] <= amount {
		return dp[amount]
	}

	return -1
}

func main() {
	testCases := []struct {
		coins  []int
		amount int
	}{
		{[]int{1, 2, 5}, 11},
		{[]int{2}, 3},
		{[]int{1}, 0},
	}

	for _, testCase := range testCases {
		fmt.Println(coinChange(testCase.coins, testCase.amount))
	}
}
