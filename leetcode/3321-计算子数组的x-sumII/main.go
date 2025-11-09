package main

import (
	"cmp"
	"fmt"

	"github.com/emirpasic/gods/v2/trees/redblacktree"
)

/*
	题目说明
	存在一个数组nums，要计算k长度的子数组中，出现次数前x大的数据的和，如果出现次数相同，先算大的。
	解题思路，通过滑动窗口+哈希表
	在长度为k的滑动窗口中，存在两类数据，一类是出现次数前x大的，一类是非前x的。
	设置L，维护前x大的数据
	设置R，维护非前x的数据
	滑动窗口移动的时候，需要加入一个元素、删除一个元素
	当加入一个元素的时候，先看看是否是前x大的数据，是则加入L，不是则加入R，如果加入L，则L中移除最小的数据
	当删除一个元素的时候，先看看L中是否有这个数据，如果有就删除，之后判断是否要换进来新的元素；如果在L中不存在就在R中，直接删除即可

	这个时候为了方便数据的维护，可以设置一个哈希，记录出现的数据和他们对应的次数

	pair(元素的出现次数， 元素值)
	pair(cnt[x], x) 可以进行双关键字排序，先排出现的次数，再排元素值，符合题目的要求
*/

type pair struct {
	cnt int
	x   int
}

func less(p, q pair) int {
	return cmp.Or(p.cnt-q.cnt, p.x-q.x)
}

func findXSum(nums []int, k, x int) []int64 {
	L := redblacktree.NewWith[pair, struct{}](less)
	R := redblacktree.NewWith[pair, struct{}](less)

	sumL := 0

	cnt := map[int]int{}

	add := func(x int) {
		p := pair{cnt[x], x}

		if p.cnt == 0 {
			return
		} // 数量为0 则不处理

		if !L.Empty() && less(p, L.Left().Key) > 0 { // p比L中最小的大，放入L中
			sumL += p.cnt * p.x
			L.Put(p, struct{}{})
		} else { // 不够大，放入R中
			R.Put(p, struct{}{})
		}
	}

	del := func(x int) {
		p := pair{cnt[x], x}

		if p.cnt == 0 {
			return
		}
		if _, ok := L.Get(p); ok { // 删除L中的数据
			sumL -= p.cnt * p.x
			L.Remove(p)
		} else { // 删除R中的数据
			R.Remove(p)
		}
	}

	l2r := func() { // L中的数据移到R中
		p := L.Left().Key
		sumL -= p.cnt * p.x
		L.Remove(p)
		R.Put(p, struct{}{})
	}

	r2l := func() { // R中的数据移到L中
		p := R.Right().Key
		sumL += p.cnt * p.x
		R.Remove(p)
		L.Put(p, struct{}{})
	}

	ans := make([]int64, len(nums)-k+1)

	for r, in := range nums {
		// 添加一个元素
		del(in)
		cnt[in] += 1
		add(in)

		l := r - k + 1
		if l < 0 { // 窗口未形成
			continue
		}

		// 维护大小
		for !R.Empty() && L.Size() < x { // L中数量不够x个
			r2l()
		}

		for L.Size() > x { // L中数量大于x个
			l2r()
		}

		ans[l] = int64(sumL)

		// 删除窗口的第一个元素
		out := nums[l]
		del(out)
		cnt[out] -= 1
		add(out)
	}

	return ans
}

func main() {
	testCases := []struct {
		nums []int
		k    int
		x    int
	}{
		{[]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2},
		{[]int{3, 8, 7, 8, 7, 5}, 2, 2},
	}

	for _, testCase := range testCases {
		fmt.Println(findXSum(testCase.nums, testCase.k, testCase.x))
	}
}
