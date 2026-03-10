package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	tmp := [256]int{}
	n := len(s)
	left := 0

	ans := 0

	for i := 0; i < n; i++ {
		cnt := int(s[i] - '0')
		if tmp[cnt] == 0 {
			// 当前数字不在窗口中
			tmp[cnt] = i + 1
		} else {
			ans = max(ans, i-left)
			for left <= tmp[cnt]-1 {
				tmp[s[left]-'0'] = 0
				left += 1
			}
			tmp[cnt] = i + 1
		}
	}
	ans = max(ans, n-left)

	return ans
}

func main() {
	testCases := []struct {
		s   string
		ans int
	}{
		{
			"abcabcbb",
			3,
		},
		{
			"bbbbb",
			1,
		},
		{
			"pwwkew",
			3,
		},
	}

	for _, testCase := range testCases {
		ans := lengthOfLongestSubstring(testCase.s)
		fmt.Println(ans, ans == testCase.ans)
	}
}
