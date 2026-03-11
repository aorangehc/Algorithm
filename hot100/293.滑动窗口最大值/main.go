package main

import (
	"fmt"
	"reflect"
)

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)

	if k > n {
		return nil
	}

	ans := make([]int, n-k+1)
	cnt := 0
	numMax := []int{} // 记录最大值的位置

	for i := 0; i < n; i++ {
		// 新的值比之前的大，之前就没用了
		// 新的值存在的时间更长
		for len(numMax) > 0 && nums[i] > nums[numMax[len(numMax)-1]] {
			numMax = numMax[:len(numMax)-1]
		}

		numMax = append(numMax, i)

		if i-numMax[0] >= k {
			// 过了最大值的位置
			numMax = numMax[1:]
		}

		if i >= k-1 {
			ans[cnt] = nums[numMax[0]]
			cnt += 1
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		k    int
		ans  []int
	}{
		{
			[]int{1, 3, -1, -3, 5, 3, 6, 7},
			3,
			[]int{3, 3, 5, 5, 6, 7},
		},
		{
			[]int{1},
			1,
			[]int{1},
		},
	}

	for _, testCase := range testCases {
		ans := maxSlidingWindow(testCase.nums, testCase.k)
		fmt.Println(ans, reflect.DeepEqual(testCase.ans, ans))
	}
}
