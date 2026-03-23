package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	// dp[i][j]表示 word1 前 i 个字符转换成 word2 前 j 个字符需要的次数
	// 初始化，dp[i][0] = i, dp[0][j] = j

	n, m := len(word1), len(word2)

	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, m)
		dp[i][0] = i
	}

	for j := range m {
		dp[0][j] = j
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}

	return dp[n-1][m-1]
}

func main() {
	testCases := []struct {
		word1 string
		word2 string
		ans   int
	}{
		{
			"horse",
			"ros",
			3,
		},
		{
			"intention",
			"execution",
			5,
		},
	}

	for _, testCase := range testCases {
		ans := minDistance(testCase.word1, testCase.word2)

		fmt.Println(ans, ans == testCase.ans)
	}
}
