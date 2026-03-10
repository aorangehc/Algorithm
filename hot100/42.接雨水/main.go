package main

import "fmt"

func trap(height []int) int {
	// 关键点是找到中间最高的
	n := len(height)
	left, right := 0, n-1
	maxLeft, maxRight := height[left], height[right]

	ans := 0

	for left < right {
		if maxLeft > maxRight {
			for left < right && height[right] <= maxRight {
				ans += maxRight - height[right]
				right -= 1
			}
			maxRight = max(maxRight, height[right])
		} else {
			for left < right && height[left] <= maxLeft {
				ans += maxLeft - height[left]
				left += 1
			}
			maxLeft = max(maxLeft, height[left])
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		height []int
		ans    int
	}{
		{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
		{
			[]int{4, 2, 0, 3, 2, 5},
			9,
		},
	}

	for _, testCase := range testCases {
		ans := trap(testCase.height)
		fmt.Println(ans, ans == testCase.ans)
	}
}
