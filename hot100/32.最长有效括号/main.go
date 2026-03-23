package main

import "fmt"

func longestValidParentheses(s string) int {
	stk := []int{-1}
	n := len(s)

	ans := 0

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stk = append(stk, i)
		} else if len(stk) > 1 {
			stk = stk[:len(stk)-1]
			ans = max(ans, i-stk[len(stk)-1])
		} else {
			stk[0] = i
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		s   string
		ans int
	}{
		{
			"(()",
			2,
		},
		{
			")()())",
			4,
		},
		{
			"",
			0,
		},
	}

	for _, testCase := range testCases {
		ans := longestValidParentheses(testCase.s)
		fmt.Println(ans, testCase.ans == ans)
	}
}
