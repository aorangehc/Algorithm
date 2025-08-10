package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := n - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left += 1
				}
				for left < right && nums[right] == nums[right-1] {
					right -= 1
				}
				left += 1
				right -= 1
			} else if sum < 0 {
				left += 1
			} else {
				right -= 1
			}
		}
	}

	return ans
}

func main() {
	testCases := [][]int{
		[]int{-1, 0, 1, 2, -1, -4},
		[]int{0, 1, 1},
		[]int{0, 0, 0},
	}

	for _, testCase := range testCases {
		ans := threeSum(testCase)
		fmt.Println(ans)
	}
}
