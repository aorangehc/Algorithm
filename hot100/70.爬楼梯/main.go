package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	a1, a2 := 1, 1

	for i := 1; i < n; i++ {
		tmp := a2
		a2 = a1 + a2
		a1 = tmp
	}

	return a2
}

func main() {
	testCases := []struct {
		n   int
		ans int
	}{
		{2, 2},
		{3, 3},
	}

	for _, testCase := range testCases {
		ans := climbStairs(testCase.n)
		fmt.Println(ans, ans == testCase.ans)
	}
}
