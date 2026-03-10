package main

import (
	"fmt"
	"reflect"
	"slices"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	slices.Sort(nums)
	ans := [][]int{}

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1

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
				for left < right && nums[left] == nums[left+1] {
					left += 1
				}
				left += 1
			} else {
				for left < right && nums[right] == nums[right-1] {
					right -= 1
				}
				right -= 1
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		ans  [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				[]int{-1, -1, 2},
				[]int{-1, 0, 1},
			},
		},
		{
			[]int{0, 1, 1},
			[][]int{},
		},
	}

	for _, testCase := range testCases {
		fmt.Println(threeSum(testCase.nums), reflect.DeepEqual(threeSum(testCase.nums), testCase.ans))
	}
}
