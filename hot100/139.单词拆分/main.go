package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	// dp[i] = dp[i - len(wordDict[j])]

	dp[0] = true

	for i := 1; i <= n; i++ {
		for _, word := range wordDict {
			lenWord := len(word)

			if lenWord <= i {
				subStr := s[i-lenWord : i]
				if subStr == word {
					if dp[i] == false {
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
		s        string
		wordDict []string
		ans      bool
	}{
		{
			"leetcode",
			[]string{
				"leet",
				"code",
			},
			true,
		},
		{
			"applepenapple",
			[]string{
				"apple",
				"pen",
			},
			true,
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
			false,
		},
	}

	for _, testCase := range testCases {
		ans := wordBreak(testCase.s, testCase.wordDict)
		fmt.Println(ans, ans == testCase.ans)
	}
}
