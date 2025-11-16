// 1176.健身计划评估
package main

import "fmt"

func dietPlanPerformance(calories []int, k, lower, upper int) int {
	ans, cnt := 0, 0

	for i, calorie := range calories {
		cnt += calorie

		if i >= k {
			cnt -= calories[i-k]
		}

		if i >= k-1 {
			if cnt > upper {
				ans += 1
			}

			if cnt < lower {
				ans -= 1
			}
		}
	}

	return ans
}

func main() {
	testCases := []struct {
		calories []int
		k        int
		lower    int
		upper    int
	}{
		{[]int{1, 2, 3, 4, 5}, 1, 3, 3},
		{[]int{3, 2}, 2, 0, 1},
		{[]int{6, 5, 0, 0}, 2, 1, 5},
	}

	for _, testCase := range testCases {
		fmt.Println(dietPlanPerformance(testCase.calories, testCase.k, testCase.lower, testCase.upper))
	}
}
