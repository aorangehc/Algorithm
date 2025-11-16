// 2904.最短且字典序最小的美丽子字符串
package main

import (
	"fmt"
	"strings"
)

func shortestBeautifulSubstring(s string, k int) string {
	if strings.Count(s, "1") < k {
		return ""
	}
	ans := s
	cnt, left := 0, 0

	for i, ch := range s {
		if ch == '1' {
			cnt += 1
		}

		for cnt >= k {
			tmp := s[left : i+1]
			if s[left] == '1' {
				cnt -= 1
			}
			left += 1
			if len(tmp) == len(ans) {
				if tmp < ans {
					ans = tmp
				}
			}
			if len(tmp) < len(ans) {
				ans = tmp
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		s string
		k int
	}{
		{"100011001", 3},
		{"1011", 2},
		{"000", 1},
	}

	for _, testCase := range testCases {
		fmt.Println(shortestBeautifulSubstring(testCase.s, testCase.k))
	}
}
