// 2461.长度为K子数组中的最大和
package main

import "fmt"

func maximumSubarraySum(nums []int, k int) int64 {
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

		if i >= k-1 && len(mp) == k {
			ans = max(ans, cnt)
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 5, 4, 2, 9, 9, 9}, 3},
		{[]int{4, 4, 4}, 3},
	}

	for _, testCase := range testCases {
		fmt.Println(maximumSubarraySum(testCase.nums, testCase.k))
	}
}
