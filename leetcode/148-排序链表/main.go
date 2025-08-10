package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := &ListNode{}
	curr := head

	for curr != nil {
		nxt := curr.Next
		pre := tmp

		for pre.Next != nil && pre.Next.Val < curr.Val {
			pre = pre.Next
		}

		curr.Next = pre.Next
		pre.Next = curr

		curr = nxt
	}

	return tmp.Next
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	cur := head

	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}

	return head
}

func printList(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Printf("%d ", p.Val)
	}

	fmt.Println()
}

func main() {
	testCases := [][]int{
		[]int{4, 2, 3, 1},
		[]int{-1, 5, 3, 4, 0},
		[]int{},
	}

	for _, testCase := range testCases {
		head := buildList(testCase)
		fmt.Println("原始链表：")
		printList(head)
		sortedList := sortList(head)
		fmt.Println("排序后链表：")
		printList(sortedList)
	}
}
