package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func buildCList(nums []int, pos int) *ListNode {
	posNode := &ListNode{}

	head := &ListNode{}
	cur := head

	for i := 0; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
		if i == pos {
			posNode = cur
		}
	}

	if pos != -1 {
		cur.Next = posNode
	}

	return head.Next
}

func main() {
	testCases := []struct {
		nums []int
		pos  int
		ans  bool
	}{
		{[]int{3, 2, 0, -4}, -1, false}, // 无环
		{[]int{3, 2, 0, -4}, 1, true},
		{[]int{1, 2}, 0, true},
		{[]int{1}, -1, false},
	}

	for _, testCase := range testCases {
		head := buildCList(testCase.nums, testCase.pos)
		fmt.Println(hasCycle(head) == testCase.ans)
	}
}
