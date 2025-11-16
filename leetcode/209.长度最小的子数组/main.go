// 209.长度最小的子数组
package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	ans, left := math.MaxInt, 0
	cnt := 0

	for i, num := range nums {
		cnt += num
		for cnt >= target {
			ans = min(ans, i-left+1)
			cnt -= nums[left]
			left += 1
		}
	}

	if ans > len(nums) {
		return 0
	} else {
		return ans
	}
}

func main() {
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 3, 1, 2, 4, 3}, 7},
		{[]int{1, 4, 4}, 4},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 11},
	}

	for _, testCase := range testCases {
		fmt.Println(minSubArrayLen(testCase.target, testCase.nums))
	}
}
