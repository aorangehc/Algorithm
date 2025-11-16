// 2379.得到K个黑块的最少涂色次数
package main

import "fmt"

func minimumRecolors(blocks string, k int) int {
	ans, cnt := 0, 0

	for i, block := range blocks {
		if block == 'B' {
			cnt += 1
		}

		if i >= k {
			if blocks[i-k] == 'B' {
				cnt -= 1
			}
		}

		if i >= k-1 {
			ans = max(ans, cnt)
		}
	}

	return k - ans
}

func main() {
	testCases := []struct {
		blocks string
		k      int
	}{
		{"WBBWWBBWBW", 7},
		{"WBWBBBW", 2},
	}

	for _, testCase := range testCases {
		fmt.Println(minimumRecolors(testCase.blocks, testCase.k))
	}
}
