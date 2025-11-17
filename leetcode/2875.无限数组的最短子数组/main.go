// 2875.无限数组的最短子数组
package main

import (
	"fmt"
	"math"
)

func minSizeSubarray(nums []int, target int) int {
	total := 0
	for _, x := range nums {
		total += x
	}

	ans := math.MaxInt
	left, sum, n := 0, 0, len(nums)
	for right := 0; right < n*2; right++ {
		sum += nums[right%n]
		for sum > target%total {
			sum -= nums[left%n]
			left++
		}
		if sum == target%total {
			ans = min(ans, right-left+1)
		}
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans + target/total*n
}

func main() {
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{1, 2, 3}, 5},
		{[]int{1, 1, 1, 2, 3}, 4},
		{[]int{2, 4, 6, 8}, 3},
	}

	for _, testCase := range testCases {
		fmt.Println(minSizeSubarray(testCase.nums, testCase.target))
	}
}
