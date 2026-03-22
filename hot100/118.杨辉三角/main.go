package main

import (
	"fmt"
	"reflect"
)

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)

	for i := range ans {
		ans[i] = make([]int, i+1)
	}

	ans[0] = []int{1}

	if numRows >= 2 {
		ans[1] = []int{1, 1}
	}

	for i := 2; i < numRows; i++ {
		ans[i][0] = 1
		ans[i][i] = 1

		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		numRows int
		ans     [][]int
	}{
		{
			5,
			[][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
				[]int{1, 3, 3, 1},
				[]int{1, 4, 6, 4, 1},
			},
		},
		{
			1,
			[][]int{
				[]int{1},
			},
		},
	}

	for _, testCase := range testCases {
		ans := generate(testCase.numRows)
		fmt.Println(ans, reflect.DeepEqual(ans, testCase.ans))
	}
}
