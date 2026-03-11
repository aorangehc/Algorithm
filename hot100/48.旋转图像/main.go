package main

import (
	"fmt"
	"reflect"
)

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}

func main() {
	testCases := []struct {
		matrix [][]int
		ans    [][]int
	}{
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[][]int{
				[]int{7, 4, 1},
				[]int{8, 5, 2},
				[]int{9, 6, 3},
			},
		},
		{
			[][]int{
				[]int{5, 1, 9, 11},
				[]int{2, 4, 8, 10},
				[]int{13, 3, 6, 7},
				[]int{15, 14, 12, 16},
			},
			[][]int{
				[]int{15, 13, 2, 5},
				[]int{14, 3, 4, 1},
				[]int{12, 6, 8, 9},
				[]int{16, 7, 10, 11},
			},
		},
	}

	for _, testCase := range testCases {
		rotate(testCase.matrix)
		fmt.Println(testCase.matrix, reflect.DeepEqual(testCase.matrix, testCase.ans))
	}
}
