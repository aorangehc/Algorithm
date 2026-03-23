package main

import "fmt"

func uniquePaths(m, n int) int {
	dp := make([][]int, m)

	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	testCases := []struct {
		m   int
		n   int
		ans int
	}{
		{3, 7, 28},
		{3, 2, 3},
		{7, 3, 28},
		{3, 3, 6},
	}

	for _, testCase := range testCases {
		ans := uniquePaths(testCase.m, testCase.n)
		fmt.Println(ans, ans == testCase.ans)
	}
}
