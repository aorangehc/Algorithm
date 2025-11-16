// 1151.最少交换次数来组合所有的1
package main

import "fmt"

func minSwaps(data []int) int {
	k := 0 // 滑动窗口的大小
	for _, val := range data {
		k += val
	}

	ans, cnt := 0, 0

	for i, val := range data {
		cnt += val

		if i >= k {
			cnt -= data[i-k]
		}

		if i >= k-1 {
			ans = max(ans, cnt)
		}
	}

	return k - ans
}

func main() {
	testCases := [][]int{
		[]int{1, 0, 1, 0, 1},
		[]int{0, 0, 0, 1, 0},
		[]int{1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1},
		[]int{1, 0, 1, 0, 1, 0, 1, 1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1},
	}

	for _, testCase := range testCases {
		fmt.Println(minSwaps(testCase))
	}
}
