// 1234.替换子串得到平衡字符串
package main

import "fmt"

func balancedString(s string) int {
	// 构建哈希表，统计各自的数量
	// 构建一个最小窗口，让窗口外的每种字符数量尽量接近n，这个最小窗口就是需要替换的字符串
	mp := make(map[byte]int, 4)
	n := len(s) / 4

	for i := 0; i < len(s); i++ {
		mp[s[i]] += 1
	}

	// 已经平衡，返回0
	if mp['Q'] == n && mp['W'] == n && mp['E'] == n && mp['R'] == n {
		return 0
	}

	ans, left := len(s), 0

	for i := 0; i < len(s); i++ {
		mp[s[i]] -= 1 // 放进窗口
		for mp['Q'] <= n && mp['W'] <= n && mp['E'] <= n && mp['R'] <= n {
			// 满足条件，记录并收缩
			ans = min(ans, i-left+1)
			mp[s[left]] += 1
			left += 1
		}
	}

	return ans
}

func main() {
	testCases := []string{
		"QWER",
		"QQWE",
		"QQQW",
		"QQQQ",
	}

	for _, testCase := range testCases {
		fmt.Println(balancedString(testCase))
	}
}
