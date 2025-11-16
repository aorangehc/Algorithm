// 2841.几乎唯一子数组的最大和
package main

import "fmt"

func maxSum(nums []int, m, k int) int64 {
	mp := make(map[int]int)
	var ans, cnt int64
	for i, num := range nums {
		cnt += int64(num)
		mp[num] += 1

		if i >= k {
			cnt -= int64(nums[i-k])
			mp[nums[i-k]] -= 1
			if mp[nums[i-k]] == 0 {
				delete(mp, nums[i-k])
			}
		}

		if i >= k-1 && len(mp) >= m {
			ans = max(ans, cnt)
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		m    int
		k    int
	}{
		{[]int{2, 6, 7, 3, 1, 7}, 3, 4},
		{[]int{5, 9, 9, 2, 4, 5, 4}, 1, 3},
		{[]int{1, 2, 1, 2, 1, 2, 1}, 3, 3},
	}

	for _, testCase := range testCases {
		fmt.Println(maxSum(testCase.nums, testCase.m, testCase.k))
	}
}
