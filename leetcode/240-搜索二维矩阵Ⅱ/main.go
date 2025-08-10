package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1

	for i < len(matrix) && j >= 0 {
		if matrix[i][j] > target {
			j -= 1
		} else if matrix[i][j] < target {
			i += 1
		} else {
			return true
		}
	}

	return false
}

func main() {
	testCases := []struct {
		matrix [][]int
		target int
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
		},
	}

	for _, testCase := range testCases {
		fmt.Println(searchMatrix(testCase.matrix, testCase.target))
	}
}
