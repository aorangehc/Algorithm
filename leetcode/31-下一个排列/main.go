package main

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums) - 1
	i := n - 1

	// 从后向前，找到第一个减少的数字
	for i >= 0 && nums[i] >= nums[i+1] {
		i -= 1
	}
	// 从后向前，找到第一个比分界大的数字，并进行交换
	if i >= 0 {
		j := n
		for j >= 0 && nums[i] >= nums[j] {
			j -= 1
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 交换分界之后的数字，变成倒序
	left, right := i+1, n
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left += 1
		right -= 1
	}
}

func main() {
	testCases := [][]int{
		[]int{1, 2, 3},
		[]int{3, 2, 1},
		[]int{1, 1, 5},
	}

	for _, testCase := range testCases {
		nextPermutation(testCase)
		fmt.Println(testCase)
	}
}
