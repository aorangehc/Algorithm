package main

import "fmt"

func longestConsecutive(nums []int) int {
	ans := 0
	tmp := map[int]bool{}

	for _, num := range nums {
		tmp[num] = true
	}

	for x := range tmp {
		if tmp[x-1] == true {
			continue
		}
		y := x + 1
		for tmp[y] {
			y += 1
		}
		ans = max(ans, y-x)
	}

	return ans
}

func main() {
	testCases := [][]int{
		[]int{100, 4, 200, 1, 3, 2},
		[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
		[]int{1, 0, 1, 2},
	}

	for _, testCase := range testCases {
		fmt.Println(longestConsecutive(testCase))
	}
}
