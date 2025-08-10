package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			for slow != head {
				head = head.Next
				slow = slow.Next
			}

			return slow
		}
	}

	return nil
}

func buildCList(nums []int, pos int) (*ListNode, *ListNode) {
	if len(nums) == 0 {
		return nil, nil
	}
	tmp := &ListNode{}
	cur := tmp
	var node *ListNode

	for i := 0; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next

		if i == pos {
			node = cur
		}
	}

	if pos != -1 {
		cur.Next = node
	}

	return tmp.Next, node
}

func main() {
	testCases := []struct {
		nums []int
		pos  int
	}{
		{[]int{3, 2, 0, -1}, 1},
		{[]int{1, 2}, 0},
		{[]int{1}, -1},
		{[]int{}, -1},
	}

	for _, testCase := range testCases {
		head, node := buildCList(testCase.nums, testCase.pos)
		findNode := detectCycle(head)

		fmt.Println(node == findNode, node)
	}
}
