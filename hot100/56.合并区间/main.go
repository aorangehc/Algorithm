package main

import (
	"fmt"
	"reflect"
	"slices"
)

func merge(intervals [][]int) [][]int {
	ans := [][]int{}

	slices.SortFunc(intervals, func(a, b []int) int { return a[0] - b[0] })

	left, right := intervals[0][0], intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > right {
			ans = append(ans, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else {
			right = max(right, intervals[i][1])
		}
	}

	ans = append(ans, []int{left, right})

	return ans
}

func main() {
	testCases := []struct {
		intervals [][]int
		ans       [][]int
	}{
		{
			[][]int{
				[]int{1, 3},
				[]int{2, 6},
				[]int{8, 10},
				[]int{15, 18},
			},
			[][]int{
				[]int{1, 6},
				[]int{8, 10},
				[]int{15, 18},
			},
		},
		{
			[][]int{
				[]int{1, 4},
				[]int{4, 5},
			},
			[][]int{
				[]int{1, 5},
			},
		},
		{
			[][]int{
				[]int{4, 7},
				[]int{1, 4},
			},
			[][]int{
				[]int{1, 7},
			},
		},
	}

	for _, testCase := range testCases {
		ans := merge(testCase.intervals)
		fmt.Println(ans, reflect.DeepEqual(ans, testCase.ans))
	}
}
