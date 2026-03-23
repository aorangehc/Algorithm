package main

import (
	"fmt"
	"sort"
)

func lengthOfLIS(nums []int) int {
	sub := []int{}

	for _, num := range nums {
		idx := sort.Search(len(sub), func(i int) bool { return sub[i] >= num })

		if idx == len(sub) {
			sub = append(sub, num)
		} else {
			sub[idx] = num
		}
	}

	return len(sub)
}

func main() {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			[]int{10, 9, 2, 5, 3, 7, 101, 18},
			4,
		},
		{
			[]int{0, 1, 0, 3, 2, 3},
			4,
		},
		{
			[]int{7, 7, 7, 7, 7, 7, 7},
			1,
		},
	}

	for _, testCase := range testCases {
		ans := lengthOfLIS(testCase.nums)
		fmt.Println(ans, ans == testCase.ans)
	}
}
