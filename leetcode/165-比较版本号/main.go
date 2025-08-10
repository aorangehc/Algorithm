package main

import (
	"fmt"
)

func compareVersion(version1, version2 string) int {
	n, m := len(version1), len(version2)
	i, j := 0, 0

	for i < n || j < m {
		num1, num2 := 0, 0

		// 读取 version1 的一段数字
		for i < n && version1[i] != '.' {
			num1 = num1*10 + int(version1[i]-'0')
			i++
		}
		// 跳过 '.'
		if i < n && version1[i] == '.' {
			i++
		}

		// 读取 version2 的一段数字
		for j < m && version2[j] != '.' {
			num2 = num2*10 + int(version2[j]-'0')
			j++
		}
		if j < m && version2[j] == '.' {
			j++
		}

		// 比较这一段
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	return 0
}

func main() {
	testCases := []struct {
		version1 string
		version2 string
	}{
		{"1.2", "1.10"},
		{"1.01", "1.001"},
		{"1.0", "1.0.0.0"},
		{"2.1", "2.0.9"},
	}

	for _, testCase := range testCases {
		fmt.Println(compareVersion(testCase.version1, testCase.version2))
	}
}
