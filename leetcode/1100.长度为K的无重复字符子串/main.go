// 1100.长度为K的无重复字符子串
package main

import "fmt"

func numKLenSubstrNoRepeats(s string, k int) int {
	mp := make(map[byte]int)
	ans := 0

	for i := 0; i < len(s); i++ {
		mp[s[i]] += 1

		if i >= k {
			mp[s[i-k]] -= 1
			if mp[s[i-k]] == 0 {
				delete(mp, s[i-k])
			}
		}

		if i >= k-1 && len(mp) == k {
			ans += 1
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		s string
		k int
	}{
		{"havefunonleetcode", 5},
		{"home", 5},
	}

	for _, testCase := range testCases {
		fmt.Println(numKLenSubstrNoRepeats(testCase.s, testCase.k))
	}
}
