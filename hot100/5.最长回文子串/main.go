package main

import "fmt"

func center(s string, left, right int) int {
	//fmt.Println(left, right)
	// 双中心，返回回文串的长度
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left -= 1
			right += 1
		} else {
			break
		}
	}

	return right - left - 1
}

func longestPalindrome(s string) string {
	left, right := 0, 0
	ans := 0

	for i := 0; i < len(s); i++ {
		len1, len2 := center(s, i, i), center(s, i, i+1)

		ans = max(ans, max(len1, len2))

		if ans > (right - left + 1) {
			left = i - (ans-1)/2
			right = i + ans/2
		}
	}

	return s[left : right+1]
}

func main() {
	testCases := []struct {
		s   string
		ans string
	}{
		{
			"aaaa",
			"aaaa",
		},
	}

	for _, testCase := range testCases {
		ans := longestPalindrome(testCase.s)
		fmt.Println(ans, ans == testCase.ans)
	}
}
