package main

import "fmt"

func judge(ch byte) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return true
	}

	return false
}

func maxVowels(s string, k int) int {
	ans := 0
	cnt := 0

	for i := 0; i < len(s); i++ {
		if judge(s[i]) {
			cnt++
		}
		if i >= k {
			if judge(s[i-k]) {
				cnt -= 1
			}
		}
		if i >= k-1 {
			ans = max(ans, cnt)
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		s string
		k int
	}{
		{"abciiidef", 3},
		{"aeiou", 2},
		{"leetcode", 3},
		{"rhythms", 4},
		{"tryhard", 4},
	}

	for _, testCase := range testCases {
		fmt.Println(maxVowels(testCase.s, testCase.k))
	}
}
