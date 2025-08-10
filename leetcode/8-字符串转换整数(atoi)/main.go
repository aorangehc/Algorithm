package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	// 主要步骤，设置最大值和最小值
	// 跳过前面的空格
	// 检查是否有符号，设置flag进行标注
	// 读取数字并检查越界
	const maxInt32 = math.MaxInt32
	const minInt32 = math.MinInt32

	n := len(s)
	idx := 0

	for idx < n && s[idx] == ' ' {
		idx += 1
	}
	if idx == n {
		return 0
	}

	sign := 1
	if s[idx] == '+' {
		idx += 1
	} else if s[idx] == '-' {
		idx += 1
		sign = -1
	}

	var ans int64 = 0
	for idx < n && s[idx] >= '0' && s[idx] <= '9' {
		digit := int64(s[idx] - '0')
		ans = ans*10 + digit
		// 越界检查
		if sign == 1 && ans > maxInt32 {
			return maxInt32
		}
		if sign == -1 && ans*-1 < minInt32 {
			return minInt32
		}
		idx += 1
	}

	return int(ans * int64(sign))
}

func main() {
	testCases := []string{
		"42",
		"-42",
		"   -42",
		"   -42     ",
		"4193 with words",
		"   +0 123",
		"   +000001234",
		"",
	}

	for _, testCase := range testCases {
		fmt.Println(myAtoi(testCase))
	}
}
