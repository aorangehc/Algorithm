package main

import "fmt"

/*
写一下本题的解题思路
设置dp数组。dp[i][j]表示word1的前i个字符转成word2[j]前j个字符需要的最小操作数
dp[i][0] = i;   dp[0][j] = j 可以进行初始化
如果word1[i] 和word2[j]相等。不需要进行额外的操作，dp[i][j] = dp[i - 1][j - 1] // 这个时候表示i和j都要从1开始
如果 word1[i] != word2[j]，可以分成三种情况
1、替换  dp[i][j] = dp[i - 1][j - 1] + 1
2、删除 dp[i][j] = dp[i - 1][j] + 1
3、插入 dp[i][j] = dp[i][j - 1]
*/

func minDistance(word1, word2 string) int {
	n, m := len(word1), len(word2)
	dp := make([][]int, n+1)

	for i := range dp {
		dp[i] = make([]int, m+1)
		dp[i][0] = i
	}

	for j := range m + 1 {
		dp[0][j] = j
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}

	return dp[n][m]

}

func main() {
	testCases := []struct {
		word1 string
		word2 string
	}{
		{"", ""},
		{"horse", ""},
		{"", "horse"},
		{"horse", "ros"},
		{"intention", "execution"},
	}

	for _, testCase := range testCases {
		ans := minDistance(testCase.word1, testCase.word2)
		fmt.Println(ans)
	}
}
