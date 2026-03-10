package main

import (
	"fmt"
	"reflect"
)

func findAnagrams(s string, p string) []int {
	countS := [26]int{}
	countP := [26]int{}

	ans := []int{}

	n, m := len(s), len(p)

	for i := 0; i < m; i++ {
		countP[p[i]-'a'] += 1
	}

	for i := 0; i < n; i++ {
		countS[s[i]-'a'] += 1

		left := i - m + 1
		if left < 0 {
			continue
		}

		if countS == countP {
			ans = append(ans, left)
		}
		countS[s[left]-'a'] -= 1
	}

	return ans
}

func main() {
	testCases := []struct {
		s   string
		p   string
		ans []int
	}{
		{
			"cbaebabacd",
			"abc",
			[]int{0, 6},
		},
		{
			"abab",
			"ab",
			[]int{0, 1, 2},
		},
	}

	for _, testCase := range testCases {
		ans := findAnagrams(testCase.s, testCase.p)
		fmt.Println(ans, reflect.DeepEqual(ans, testCase.ans))
	}
}
