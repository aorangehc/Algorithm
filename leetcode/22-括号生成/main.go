package main

import "fmt"

func generateParenthesis(n int) []string {
	m := n * 2

	path := make([]byte, m)
	ans := []string{}

	// i 有多少括号
	// left 左括号的数量

	var dfs func(int, int)

	dfs = func(i, left int) {
		if i == m {
			ans = append(ans, string(path))

			return
		}

		if left < n {
			path[i] = '('
			dfs(i+1, left+1)
		}
		if i-left < left {
			path[i] = ')'
			dfs(i+1, left)
		}
	}

	dfs(0, 0)

	return ans
}

func main() {
	for i := range 5 {
		ans := generateParenthesis(i)
		fmt.Println(ans)
	}
}
