// 1852.每个子数组的数字种类数
package main

import "fmt"

func distinctNumbers(nums []int, k int) []int {
	mp := make(map[int]int)
	ans := make([]int, len(nums)-k+1)

	for i, num := range nums {
		mp[num] += 1

		if i >= k {
			mp[nums[i-k]] -= 1
			if mp[nums[i-k]] == 0 {
				delete(mp, nums[i-k])
			}
		}

		if i >= k-1 {
			ans[i-k+1] = len(mp)
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 2, 3, 2, 2, 1, 3}, 3},
		{[]int{1, 1, 1, 1, 2, 3, 4}, 4},
	}

	for _, testCase := range testCases {
		fmt.Println(distinctNumbers(testCase.nums, testCase.k))
	}
}
