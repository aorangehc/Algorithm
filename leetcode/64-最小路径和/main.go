package main

import "fmt"

func minPathSum(grid [][]int) int {
	n, m := len(grid), len(grid[0])

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < m; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[n-1][m-1]
}

func main() {
	testCases := [][][]int{
		[][]int{
			[]int{1, 3, 1},
			[]int{1, 5, 1},
			[]int{4, 2, 1},
		},
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
		},
	}
	for _, testCase := range testCases {
		fmt.Println(minPathSum(testCase))
	}
}
