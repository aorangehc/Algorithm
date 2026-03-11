package main

import "fmt"

func maxSubArray(nums []int) int {
	ans := nums[0]
	cnt := nums[0]

	for i := 1; i < len(nums); i++ {
		cnt = max(cnt+nums[i], nums[i])
		ans = max(ans, cnt)
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{5, 4, -1, 7, 8},
			23,
		},
	}

	for _, testCase := range testCases {
		ans := maxSubArray(testCase.nums)
		fmt.Println(ans, ans == testCase.ans)
	}
}
