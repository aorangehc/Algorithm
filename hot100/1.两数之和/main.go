package main

import (
	"fmt"
	"reflect"
)

func twoSum(nums []int, target int) (ans []int) {
	mp := map[int]int{}

	for idx, num := range nums {
		if i, ok := mp[target-num]; ok {
			ans = []int{i, idx}
		}
		mp[num] = idx
	}

	return ans
}

func main() {
	testCases := []struct {
		nums   []int
		target int
		ans    []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, ans: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, ans: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, ans: []int{0, 1}},
	}

	for _, t := range testCases {
		fmt.Println(reflect.DeepEqual(twoSum(t.nums, t.target), t.ans))
	}
}
