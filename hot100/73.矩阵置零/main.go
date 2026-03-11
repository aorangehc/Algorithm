package main

import (
	"fmt"
	"reflect"
)

func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])

	row := make([]bool, n)
	col := make([]bool, m)

	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] {
				r[j] = 0
			}
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
				[]int{1, 1, 1},
				[]int{1, 0, 1},
				[]int{1, 1, 1},
			},
			[][]int{
				[]int{1, 0, 1},
				[]int{0, 0, 0},
				[]int{1, 0, 1},
			},
		},
		{
			[][]int{
				[]int{0, 1, 2, 0},
				[]int{3, 4, 5, 2},
				[]int{1, 3, 1, 5},
			},
			[][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 4, 5, 0},
				[]int{0, 3, 1, 0},
			},
		},
	}

	for _, testCase := range testCases {
		setZeroes(testCase.matrix)
		fmt.Println(testCase.matrix, reflect.DeepEqual(testCase.ans, testCase.matrix))
	}
}
