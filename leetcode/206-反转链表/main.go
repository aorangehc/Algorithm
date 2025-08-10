// leetcode-206-翻转链表
package main

import (
	"fmt"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

// 翻转链表采用插入法
// 创建一个空节点，然后一个个插入
func reverseList(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head

	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}

func buildList(nums []int) *ListNode {
	head := &ListNode{}
	cur := head

	for _, v := range nums {
		cur.Next = &ListNode{Value: v}
		cur = cur.Next
	}

	return head.Next
}

func printList(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Printf("%d ", p.Value)
	}
	fmt.Println()
}
func main() {
	testCase := [][]int{
		{1, 2, 3, 4, 5},
		{1},
		{2, 1},
		{},
	}
	for _, c := range testCase {
		head := buildList(c)
		reversed := reverseList(head)
		printList(reversed)
	}
}
