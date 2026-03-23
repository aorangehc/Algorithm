package main

import "fmt"

func maxProduct(nums []int) int {
	// 可能存在负数，存储一个最大值和最小值
	n := len(nums)

	ans, maxNum, minNum := nums[0], nums[0], nums[0]

	for i := 1; i < n; i++ {
		a, b := maxNum*nums[i], minNum*nums[i]
		maxNum = max(nums[i], max(a, b))
		minNum = min(nums[i], min(a, b))

		ans = max(ans, maxNum)
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{
			[]int{2, 3, -2, 4},
			6,
		},
		{
			[]int{-2, 0, -1},
			0,
		},
	}

	for _, testCase := range testCases {
		ans := maxProduct(testCase.nums)
		fmt.Println(ans, ans == testCase.ans)
	}
}
