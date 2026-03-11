package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	i, j := 0, m-1
	for i < n && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			j -= 1
		} else {
			i += 1
		}
	}

	return false
}

func main() {
	testCases := []struct {
		matrix [][]int
		target int
		ans    bool
	}{
		{
			[][]int{
				[]int{1, 4, 7, 11, 15},
				[]int{2, 5, 8, 12, 19},
				[]int{3, 6, 9, 16, 22},
				[]int{10, 13, 14, 17, 24},
				[]int{18, 21, 23, 26, 30},
			},
			5,
			true,
		},
		{
			[][]int{
				[]int{1, 4, 7, 11, 15},
				[]int{2, 5, 8, 12, 19},
				[]int{3, 6, 9, 16, 22},
				[]int{10, 13, 14, 17, 24},
				[]int{18, 21, 23, 26, 30},
			},
			20,
			false,
		},
		{
			[][]int{
				[]int{-5},
			},
			-5,
			true,
		},
	}

	for _, testCase := range testCases {
		ans := searchMatrix(testCase.matrix, testCase.target)
		fmt.Println(ans == testCase.ans)
	}
}
