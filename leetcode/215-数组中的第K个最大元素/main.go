package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || k < 1 || k > len(nums) {
		return -111
	}
	n := len(nums)

	return quickSelect(nums, 0, n-1, n-k) // 查找第k大元素
}

func quickSelect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}

	partition := nums[l]
	i := l - 1
	j := r + 1

	for i < j {
		for i += 1; nums[i] < partition; i++ {
		}
		for j -= 1; nums[j] > partition; j-- {
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	if j >= k {
		return quickSelect(nums, l, j, k)
	} else {
		return quickSelect(nums, j+1, r, k)
	}
}
func main() {
	testNums := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 2},
		{[]int{6, 5, 4, 3, 2, 1}, 3},
		{[]int{1, 1, 1, 2, 2}, 4},
		{[]int{5}, 2},
	}

	for _, testnum := range testNums {
		ans := findKthLargest(testnum.nums, testnum.k)
		fmt.Println(ans)
	}
}
