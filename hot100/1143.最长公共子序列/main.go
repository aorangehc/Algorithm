package main

import "fmt"

func longestCommonSubsequence(text1, text2 string) int {
	n, m := len(text1), len(text2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i, x := range text1 {
		for j, y := range text2 {
			if x == y {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[n][m]
}

func main() {
	testCases := []struct {
		text1 string
		text2 string
		ans   int
	}{
		{
			"abcde",
			"ace",
			3,
		},
		{
			"abc",
			"abc",
			3,
		},
		{
			"abc",
			"def",
			0,
		},
	}

	for _, testCase := range testCases {
		ans := longestCommonSubsequence(testCase.text1, testCase.text2)

		fmt.Println(ans, ans == testCase.ans)
	}
}
