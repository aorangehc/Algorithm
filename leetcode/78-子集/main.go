package main

import "fmt"

func subsets(nums []int) [][]int {
	ans := [][]int{}
	n := len(nums)

	tmp := []int{}

	var dfs func(int)

	dfs = func(i int) {
		ans = append(ans, append([]int(nil), tmp...))

		if i == n {
			return
		}

		for j := i; j < n; j++ {
			tmp = append(tmp, nums[j])
			dfs(j + 1)
			tmp = tmp[:len(tmp)-1]
		}

		return
	}

	dfs(0)

	return ans
}

func main() {
	testCases := [][]int{
		[]int{1, 2, 3},
		[]int{0},
		[]int{},
	}

	for _, testCase := range testCases {
		ans := subsets(testCase)
		fmt.Println(ans)
	}
}
