package main

import "fmt"

func subarraySum(nums []int, k int) int {
	// 构建前缀和数组
	// 将当前的数字作为要统计的最后一个数字
	n := len(nums)
	preNums := make([]int, n+1)

	for i := 1; i < n+1; i++ {
		preNums[i] = preNums[i-1] + nums[i-1]
	}

	// 构建哈希表，统计出现的前缀和数量
	mp := map[int]int{}

	ans := 0

	for i := 0; i < n+1; i++ {
		// fmt.Println(preNums[i])

		ans += mp[preNums[i]-k]

		mp[preNums[i]] += 1
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		k    int
		ans  int
	}{
		{
			[]int{1, 1, 1},
			2,
			2,
		},
		{
			[]int{1, 2, 3},
			3,
			2,
		},
		{
			[]int{1},
			0,
			0,
		},
	}

	for _, testCase := range testCases {
		ans := subarraySum(testCase.nums, testCase.k)
		fmt.Println(ans, ans == testCase.ans)
	}
}
