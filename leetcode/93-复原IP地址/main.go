package main

import (
	"fmt"
	"strconv"
)

func isValid(s string) bool {
	if len(s) == 0 || len(s) > 3 {
		return false
	}

	if s[0] == '0' && len(s) > 1 {
		return false
	}

	num, err := strconv.Atoi(s)

	if err != nil || num > 255 {
		return false
	}

	return true
}

func result(s string) []string {
	ans := []string{}

	var f func(start int, path []string)

	f = func(start int, path []string) {
		if len(path) == 4 {
			if start == len(s) {
				allPath := path[0] + "." + path[1] + "." + path[2] + "." + path[3]
				ans = append(ans, allPath)
			}

			return
		}

		for i := 1; i <= 3; i++ {
			if start+i > len(s) {
				break
			}
			segment := s[start : start+i]

			if isValid(segment) {
				f(start+i, append(path, segment))
			}
		}
	}

	f(0, []string{})

	return ans
}

func main() {
	testCases := []string{
		"25525511135",
		"0000",
		"101023",
	}

	for _, testCase := range testCases {
		ans := result(testCase)
		fmt.Println(ans)
	}
}
