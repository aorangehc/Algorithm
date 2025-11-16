// 1423.可获得的最大点数
package main

import "fmt"

func maxScore(cardPoints []int, k int) int {
	left := k - 1
	ans, cnt := 0, 0

	for i := 0; i < k; i++ {
		cnt += cardPoints[i]
	}

	ans = cnt

	for i := len(cardPoints) - 1; i >= len(cardPoints)-k; i-- {
		cnt -= cardPoints[left]
		left -= 1
		cnt += cardPoints[i]

		ans = max(ans, cnt)
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 1}, 3},
		{[]int{2, 2, 2}, 2},
		{[]int{9, 7, 7, 9, 7, 7, 9}, 7},
		{[]int{1, 1000, 1}, 1},
		{[]int{1, 79, 80, 1, 1, 1, 200, 1}, 3},
	}

	for _, testCase := range testCases {
		fmt.Println(maxScore(testCase.nums, testCase.k))
	}
}
