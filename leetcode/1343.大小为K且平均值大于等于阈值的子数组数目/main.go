// 1343.大小为K且平均值大于等于阈值的子数组数目
package main

import "fmt"

func numOfSubarrays(arr []int, k, threshold int) int {
	ans, cnt := 0, 0

	for i, num := range arr {
		cnt += num
		if i >= k {
			cnt -= arr[i-k]
		}

		if i >= k-1 && cnt >= k*threshold {
			ans += 1
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		arr       []int
		k         int
		threshold int
	}{
		{[]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4},
		{[]int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, 3, 5},
	}

	for _, testCase := range testCases {
		fmt.Println(numOfSubarrays(testCase.arr, testCase.k, testCase.threshold))
	}
}
