package main

import "fmt"

func minWindow(s, t string) string {
	// 构建数组，统计数量
	countT := [256]int{}
	cnt := 0

	for i := 0; i < len(t); i++ {
		if countT[t[i]-'0'] == 0 {
			cnt += 1
		}
		countT[t[i]-'0'] += 1
	}

	// 记录一下答案的其实是 index
	ansLeft, ansRight := -1, len(s)
	left := 0

	for i := 0; i < len(s); i++ {
		countT[s[i]-'0'] -= 1
		if countT[s[i]-'0'] == 0 {
			cnt -= 1
		}

		for cnt == 0 {
			// 更新值
			if i-left < ansRight-ansLeft {
				ansLeft, ansRight = left, i
			}

			x := s[left] - '0'

			if countT[x] == 0 {
				cnt += 1
			}

			countT[x] += 1
			left += 1
		}
	}

	if ansLeft < 0 {
		return ""
	}

	return s[ansLeft : ansRight+1]
}

func main() {
	testCases := []struct {
		s   string
		t   string
		ans string
	}{
		{
			"ADOBECODEBANC",
			"ABC",
			"BANC",
		},
		{
			"a",
			"a",
			"a",
		},
		{
			"a",
			"aa",
			"",
		},
	}

	for _, testCase := range testCases {
		ans := minWindow(testCase.s, testCase.t)
		fmt.Println(ans, ans == testCase.ans)
	}
}
