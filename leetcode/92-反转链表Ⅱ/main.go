package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left, right int) *ListNode {
	tmp := &ListNode{Next: head}

	pre := tmp
	cur := head

	for i := 1; i < left; i++ {
		pre = pre.Next
		cur = cur.Next
	}

	for i := left; i < right; i++ {
		nxt := cur.Next
		cur.Next = nxt.Next
		nxt.Next = pre.Next
		pre.Next = nxt
	}

	return tmp.Next
}

func buildList(nums []int) *ListNode {
	head := &ListNode{}
	cur := head

	for i := 0; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}

	return head.Next
}

func printList(head *ListNode) {
	for node := head; node != nil; node = node.Next {
		fmt.Printf("%d ", node.Val)
	}
	fmt.Println()
}

func main() {
	testCases := []struct {
		nums  []int
		left  int
		right int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4},
		{[]int{5, 1}, 1, 1},
		{[]int{5, 1}, 1, 2},
	}

	for i, testCase := range testCases {
		root := buildList(testCase.nums)
		fmt.Println("原始链表", i+1)
		printList(root)
		reversed := reverseBetween(root, testCase.left, testCase.right)
		fmt.Println("翻转后的链表", i+1)
		printList(reversed)
	}
}
