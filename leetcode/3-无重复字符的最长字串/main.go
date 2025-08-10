package main

import (
	"fmt"
)

// 设置left，表示当前滑动窗口的起始位置
// 设置map存储字符和最新出现的位置
// 遍历字符串，判断当前字符是否在map中，如果没有加进去，如果有，更新滑动窗口的起始位置
// 更新ans，选择已有的ans或者当前滑动窗口大小中大的一个
func lengthOfLongestSubstring(s string) int {
	ans := 0
	left := -1
	mp := make(map[rune]int)

	for i, c := range s {
		if index, ok := mp[c]; ok {
			left = max(left, index)
		}
		mp[c] = i
		ans = max(ans, i-left)
	}

	return ans
}

func main() {
	var s string
	fmt.Scanf("%s", &s)

	num := lengthOfLongestSubstring(s)

	fmt.Println(num)

	fmt.Println(lengthOfLongestSubstring("bbbbbb") == 1)
	fmt.Println(lengthOfLongestSubstring("abcabcbb") == 3)
	fmt.Println(lengthOfLongestSubstring("abcdefgbcdefghij") == 9)
}
