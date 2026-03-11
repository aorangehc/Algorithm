package main

import (
	"fmt"
	"reflect"
)

func spiralOrder(matrix [][]int) []int {
	n, m := len(matrix), len(matrix[0])
	// 记录四个角
	x, y, z, k := 0, n-1, 0, m-1

	ans := make([]int, n*m)

	cnt := 0

	for cnt < n*m {
		for i := z; i <= k; i++ {
			ans[cnt] = matrix[x][i]
			cnt += 1
		}
		x += 1
		if x > y {
			break
		}

		for i := x; i <= y; i++ {
			ans[cnt] = matrix[i][k]
			cnt += 1
		}
		k -= 1
		if z > k {
			break
		}

		for i := k; i >= z; i-- {
			ans[cnt] = matrix[y][i]
			cnt += 1
		}
		y -= 1

		if x > y {
			break
		}

		for i := y; i >= x; i-- {
			ans[cnt] = matrix[i][z]
			cnt += 1
		}

		z += 1
		if z > k {
			break
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		matrix [][]int
		ans    []int
	}{
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{
				[]int{1, 2, 3, 4},
				[]int{5, 6, 7, 8},
				[]int{9, 10, 11, 12},
			},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}

	for _, testCase := range testCases {
		ans := spiralOrder(testCase.matrix)
		fmt.Println(ans, reflect.DeepEqual(ans, testCase.ans))
	}
}
