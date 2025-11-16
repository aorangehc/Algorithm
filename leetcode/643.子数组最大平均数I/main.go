// 643.子数组最大平均数I
package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	ans, cnt := math.MinInt, 0

	for i, num := range nums {
		cnt += num

		if i >= k {
			cnt -= nums[i-k]
		}

		if i >= k-1 {
			ans = max(ans, cnt)
		}
	}

	return float64(ans) / float64(k)
}

func main() {
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 12, -5, -6, 50, 3}, 4},
		{[]int{5}, 1},
	}

	for _, testCase := range testCases {
		fmt.Println(findMaxAverage(testCase.nums, testCase.k))
	}
}
