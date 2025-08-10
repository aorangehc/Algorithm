package main

import "fmt"

func result(text1, text2 string) int {
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
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[n][m]
}

func main() {
	testCases := []struct {
		text1 string
		text2 string
	}{
		{"abcde", "ace"},
		{"abc", "abc"},
		{"abc", "def"},
	}

	for _, testCase := range testCases {
		ans := result(testCase.text1, testCase.text2)
		fmt.Println(ans, testCase.text1, testCase.text2)
	}
}
