// 76.最小覆盖子串
package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	// 先统计一下t中各个字符的数量,统计一下字符的种类
	mp := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		mp[t[i]] += 1
	}
	left := 0
	cnt := len(mp) // 字符的种类
	ans := ""
	tmp := math.MaxInt

	for i := 0; i < len(s); i++ {
		if _, ok := mp[s[i]]; ok { // 字符有效
			mp[s[i]] -= 1
			if mp[s[i]] == 0 { // 符合要求
				cnt -= 1
			}
		}

		for cnt == 0 { // 子串中已经满足要求，缩小子串范围
			if tmp > i-left+1 {
				tmp = i - left + 1
				ans = s[left : i+1]
			}
			if _, ok := mp[s[left]]; ok { // 字符有效
				mp[s[left]] += 1
				if mp[s[left]] == 1 { // 符合要求
					cnt += 1
				}
			}
			left += 1
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		s string
		t string
	}{
		{"ADOBECODEBANC", "ABC"},
		{"a", "a"},
		{"a", "aa"},
	}

	for _, testCase := range testCases {
		fmt.Println(minWindow(testCase.s, testCase.t))
	}
}
