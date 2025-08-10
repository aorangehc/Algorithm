package main

import "fmt"

func rotate(matrix [][]int) {
	n := len(matrix)

	tmp := make([][]int, n)

	for i := range tmp {
		tmp[i] = make([]int, n)
	}

	for i := range n {
		for j := range n {
			tmp[i][j] = matrix[i][j]
		}
	}

	for i := range n {
		for j := range n {
			matrix[j][n-i-1] = tmp[i][j]
		}
	}
}

func main() {
	testCases := [][][]int{
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		},
		[][]int{
			[]int{1, 2, 3, 4},
			[]int{5, 6, 7, 8},
			[]int{9, 10, 11, 12},
			[]int{13, 14, 15, 16},
		},
	}

	for _, testCase := range testCases {
		rotate(testCase)
		fmt.Println(testCase)
	}
}
