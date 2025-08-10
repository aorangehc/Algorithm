package main

import "fmt"

// 找一下left和right
// left是大于等于target
// right是小于等于target
// 确定有等于，否则使用-1

func findLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}

func findRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if right >= 0 && nums[right] == target {
		return right
	}

	return -1
}

func searchRange(nums []int, target int) []int {
	ans := make([]int, 2)
	ans[0], ans[1] = -1, -1

	left := findLeft(nums, target)
	if left == -1 {
		return ans
	}
	ans[0] = left

	ans[1] = findRight(nums, target)

	return ans
}

func main() {
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8},
		{[]int{5, 7, 7, 8, 8, 10}, 6},
		{[]int{}, 0},
	}

	for _, testCase := range testCases {
		ans := searchRange(testCase.nums, testCase.target)
		fmt.Println(ans)
	}
}
