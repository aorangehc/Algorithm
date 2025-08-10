package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for _, word := range wordDict {
			lenWord := len(word)
			if lenWord <= i {
				subStr := s[i-lenWord : i] // 截出相同长度子串
				if subStr == word {
					if dp[i] == false { // 当前是false，有机会更新成true
						dp[i] = dp[i-lenWord]
					}
				}
			}
		}
	}
	return dp[n]
}

func main() {
	testCases := []struct {
		word     string
		wordDict []string
	}{
		{
			"leetcode",
			[]string{
				"leet",
				"code",
			},
		},
		{
			"applepenapple",
			[]string{
				"apple",
				"pen",
			},
		},
		{
			"catsandog",
			[]string{
				"cats",
				"dog",
				"sand",
				"and",
				"cat",
			},
		},
	}

	for _, testCase := range testCases {
		fmt.Println(wordBreak(testCase.word, testCase.wordDict))
	}
}
