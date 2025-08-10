package main

import (
	"fmt"
	"strings"
)

func multiply(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n, m := len(num1)-1, len(num2)-1
	ans := make([]int, m+n+2)

	for i := n; i >= 0; i-- {
		for j := m; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			sum := mul + ans[i+j+1]
			ans[i+j+1] = sum % 10
			ans[i+j] += sum / 10
		}
	}
	// 去掉前导 0
	var sb strings.Builder
	i := 0
	for i < len(ans) && ans[i] == 0 {
		i++
	}
	for ; i < len(ans); i++ {
		sb.WriteByte(byte(ans[i] + '0'))
	}
	return sb.String()
}

func main() {
	testCases := []struct {
		a string
		b string
	}{
		{"1", "2"},
		{"2", "3"},
		{"123", "345"},
	}

	for _, testCase := range testCases {
		ans := multiply(testCase.a, testCase.b)
		fmt.Println(ans)
	}

}
