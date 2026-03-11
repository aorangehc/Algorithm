package main

import "fmt"

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] >= 1 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if i != nums[i]-1 {
			return i + 1
		}
	}

	return n + 1
}

func main() {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			[]int{1, 2, 0},
			3,
		},
		{
			[]int{3, 4, -1, 1},
			2,
		},
		{
			[]int{7, 8, 9, 11, 12},
			1,
		},
	}

	for _, testCase := range testCases {
		ans := firstMissingPositive((testCase.nums))
		fmt.Println(ans, ans == testCase.ans)
	}
}
