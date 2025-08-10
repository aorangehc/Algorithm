package main

import (
	"fmt"
	"slices"
)

func merge(nums [][]int) [][]int {
	ans := [][]int{}

	slices.SortFunc(nums, func(p, q []int) int { return p[0] - q[0] })
	left, right := nums[0][0], nums[0][1]

	for i := 1; i < len(nums); i++ {
		if nums[i][0] <= right {
			right = max(right, nums[i][1])
		} else {
			ans = append(ans, []int{left, right})
			left, right = nums[i][0], nums[i][1]
		}
	}

	ans = append(ans, []int{left, right})

	return ans
}

func main() {
	testCases := [][][]int{
		[][]int{
			[]int{1, 3},
			[]int{2, 6},
			[]int{8, 10},
			[]int{15, 18},
		},
		[][]int{
			[]int{1, 4},
			[]int{4, 5},
		},
	}

	for _, testCase := range testCases {
		ans := merge(testCase)
		fmt.Println(ans)

	}
}
