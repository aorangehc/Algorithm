// 2090.半径为k的子数组平均值
package main

import "fmt"

func getAverages(nums []int, k int) []int {
	ans := make([]int, len(nums))

	for i := range ans {
		ans[i] = -1
	}

	cnt := 0

	for i, num := range nums {
		cnt += num

		if i >= 2*k {
			ans[i-k] = cnt / (k*2 + 1)
			cnt -= nums[i-2*k]
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3},
		{[]int{100000}, 0},
		{[]int{8}, 100000},
	}

	for _, testCase := range testCases {
		fmt.Println(getAverages(testCase.nums, testCase.k))
	}
}
