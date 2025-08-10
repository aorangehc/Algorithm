package main

import "fmt"

func twoSum(nums []int, target int) []int {
	tmp := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if p, ok := tmp[target-nums[i]]; ok {
			return []int{p, i}
		}
		tmp[nums[i]] = i
	}

	return nil
}

func main() {
	testCases := []struct {
		nums   []int
		target int
	}{
		{[]int{2, 7, 11, 15}, 9},
		{[]int{3, 2, 4}, 6},
		{[]int{3, 3}, 6},
	}

	for _, testCase := range testCases {
		ans := twoSum(testCase.nums, testCase.target)
		fmt.Println(ans)
	}
}
