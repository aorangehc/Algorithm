// 2107.分享K个糖果后独特口味的数量
package main

import "fmt"

func shareCandies(candies []int, k int) int {
	mp := make(map[int]int)
	ans := 0

	for _, candie := range candies {
		mp[candie] += 1
	}

	for i, candie := range candies {
		mp[candie] -= 1

		if mp[candie] == 0 {
			delete(mp, candie)
		}

		if i >= k {
			mp[candies[i-k]] += 1
		}

		if i >= k-1 {
			ans = max(ans, len(mp))
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		candies []int
		k       int
	}{
		{[]int{1, 2, 2, 3, 4, 3}, 3},
		{[]int{2, 2, 2, 2, 3, 3}, 2},
		{[]int{2, 4, 5}, 0},
	}

	for _, testCase := range testCases {
		fmt.Println(shareCandies(testCase.candies, testCase.k))
	}
}
