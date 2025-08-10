package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1, num2 string) string {
	l1, l2 := len(num1)-1, len(num2)-1
	pre := 0
	var ans string

	for l1 >= 0 && l2 >= 0 {
		num := int(num1[l1]) - int('0') + int(num2[l2]) - int('0') + pre
		pre = num / 10
		num %= 10
		ans = strconv.Itoa(num) + ans
		l1 -= 1
		l2 -= 1
	}

	for l1 >= 0 {
		num := int(num1[l1]) - int('0') + pre
		pre = num / 10
		num %= 10
		ans = strconv.Itoa(num) + ans
		l1 -= 1
	}

	for l2 >= 0 {
		num := int(num2[l2]) - int('0') + pre
		pre = num / 10
		num %= 10
		ans = strconv.Itoa(num) + ans
		l2 -= 1
	}

	if pre > 0 {
		ans = strconv.Itoa(pre) + ans
	}

	return ans
}

func main() {
	testCases := []struct {
		num1 string
		num2 string
	}{
		{"11", "123"},
		{"456", "77"},
		{"0", "0"},
		{"1", "0"},
		{"0", "1"},
	}

	for _, testCase := range testCases {
		ans := addStrings(testCase.num1, testCase.num2)
		fmt.Println(ans)
	}
}
