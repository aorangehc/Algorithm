package main

import "fmt"

func mySqrt(x int) int {
	// 使用二分查找
	left, right := 0, x
	for left <= right {
		mid := (left + right) / 2
		if mid*mid <= x && (mid+1)*(mid+1) > x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return 0
}

func main() {
	for i := range 99 {
		fmt.Println(mySqrt(i))
	}
}
