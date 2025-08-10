package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(a, b []int) float64 {
	if len(a) > len(b) {
		a, b = b, a
	}

	m, n := len(a), len(b)

	a = append([]int{math.MinInt}, append(a, math.MaxInt)...)
	b = append([]int{math.MinInt}, append(b, math.MaxInt)...)

	i, j := 0, (m+n+1)/2

	for a[i+1] <= b[j] {
		i++
		j--
	}

	max1 := max(a[i], b[i])
	min2 := min(a[i+1], b[j+1])

	if (m+n)%2 > 0 {
		return float64(max1)
	}

	return float64(max1+min2) / 2
}

func main() {
	testCases := []struct {
		nums1 []int
		nums2 []int
	}{
		{[]int{1, 3}, []int{2}},
		{[]int{1, 2}, []int{3, 4}},
	}

	for _, testCase := range testCases {
		ans := findMedianSortedArrays(testCase.nums1, testCase.nums2)
		fmt.Println(ans)
	}
}
