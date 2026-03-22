package main

import "fmt"

func numSquares(n int) int {
	dp := make([]int, n+1)

	// dp[i] = min(dp[i], dp[i - j * j] + 1)

	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}

	return dp[n]
}

func main() {
	testCases := []struct {
		n   int
		ans int
	}{
		{12, 3},
		{13, 2},
	}

	for _, testCase := range testCases {
		ans := numSquares(testCase.n)

		fmt.Println(ans, ans == testCase.ans)
	}
}
