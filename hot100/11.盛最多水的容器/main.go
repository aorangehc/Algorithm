package main

import "fmt"

func maxArea(height []int) int {
	n := len(height)
	left, right := 0, n-1

	ans := 0

	for left < right {
		ans = max(ans, min(height[left], height[right])*(right-left))
		if height[left] > height[right] {
			right -= 1
		} else {
			left += 1
		}
	}
	return ans
}

func main() {
	testCases := []struct {
		height []int
		ans    int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for _, testCase := range testCases {
		fmt.Println(maxArea(testCase.height), maxArea(testCase.height) == testCase.ans)
	}
}
