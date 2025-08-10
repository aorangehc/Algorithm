package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	tmp := &ListNode{Next: head}
	pre := tmp
	cur := head

	for cur != nil && cur.Next != nil {
		curNext := cur.Next
		curNext2 := cur.Next.Next

		cur.Next = curNext2
		curNext.Next = cur

		pre.Next = curNext

		pre = cur
		cur = cur.Next
	}

	return tmp.Next
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	p := head

	for i := 1; i < len(nums); i++ {
		p.Next = &ListNode{Val: nums[i]}
		p = p.Next
	}
	return head
}

func printList(head *ListNode) {
	for node := head; node != nil; node = node.Next {
		fmt.Printf("%d ", node.Val)
	}
	fmt.Println()
}

func main() {
	testCases := [][]int{
		[]int{1, 2, 3, 4, 5, 6},
		[]int{6, 5, 4, 3, 2, 1},
		[]int{1, 2, 3},
		[]int{1},
		[]int{},
	}

	for _, testCase := range testCases {
		head := buildList(testCase)
		fmt.Println("交换之前:")
		printList(head)
		swapList := swapPairs(head)
		fmt.Println("交换之后:")
		printList(swapList)
	}
}
