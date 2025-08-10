package main

import (
	"fmt"
	"slices"
)

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	n := len(candidates)
	slices.Sort(candidates)

	path := []int{}

	var dfs func(int, int)
	dfs = func(i, left int) {
		if left == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		if i >= n || left < candidates[i] {
			return
		}
		// 不选当前的数字
		dfs(i+1, left)
		path = append(path, candidates[i])
		// 选当前的数字
		dfs(i, left-candidates[i])
		path = path[:len(path)-1]
	}
	dfs(0, target)

	return ans
}

func main() {
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 3, 6, 7}, 7},
		{[]int{2, 3, 5}, 8},
		{[]int{2}, 1},
	}

	for _, testCase := range testCases {
		fmt.Println(combinationSum(testCase.nums, testCase.target))
	}
}
