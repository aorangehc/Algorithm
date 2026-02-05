package main

import "fmt"

func longestConsecutive(nums []int) int {
	ans := 0
	mp := map[int]bool{}

	for _, num := range nums {
		mp[num] = true
	}

	for i := range mp {
		if mp[i-1] == true {
			continue
		}
		cnt := 1

		for mp[i+1] == true {
			i += 1
			cnt += 1
		}

		ans = max(ans, cnt)
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		ans  int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 0, 1, 2}, 3},
	}

	for _, t := range testCases {
		fmt.Println(longestConsecutive(t.nums) == t.ans)
	}
}
