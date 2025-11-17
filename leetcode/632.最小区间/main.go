// 632.最小区间
package main

import (
	"fmt"
	"slices"
)

func smallestRange(nums [][]int) []int {
	type pair struct{ x, i int } // 数字和他们所在的位置
	pairs := []pair{}

	for i, arr := range nums {
		for _, num := range arr {
			pairs = append(pairs, pair{num, i})
		}
	}

	slices.SortFunc(pairs, func(a, b pair) int { return a.x - b.x })

	ansL, ansR := pairs[0].x, pairs[len(pairs)-1].x
	mp := make(map[int]int)
	left := 0

	for _, p := range pairs {
		r, i := p.x, p.i
		mp[i] += 1

		for len(mp) == len(nums) { // 每个列表都至少包含一个数
			l, i := pairs[left].x, pairs[left].i
			if r-l < ansR-ansL {
				ansL, ansR = l, r
			}
			mp[i] -= 1
			if mp[i] == 0 {
				delete(mp, i)
			}
			left += 1
		}
	}
	return []int{ansL, ansR}
}

func main() {
	testCases := [][][]int{
		[][]int{
			[]int{4, 10, 15, 24, 26},
			[]int{0, 9, 12, 20},
			[]int{5, 18, 22, 30},
		},
		[][]int{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
	}

	for _, testCase := range testCases {
		fmt.Println(smallestRange(testCase))
	}
}
