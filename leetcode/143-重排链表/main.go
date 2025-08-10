package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 返回中间节点
func middleList(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur != nil {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}

	return pre
}

// 合并前后两个两个链表
func mergeList(list1, list2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head

	for list1 != nil && list2 != nil {
		p.Next = list1
		list1 = list1.Next
		p = p.Next

		p.Next = list2
		list2 = list2.Next
		p = p.Next
	}

	if list1 != nil {
		p.Next = list1
	}

	return head.Next
}

// 重排链表
func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	mid := middleList(head)
	list1 := head
	list2 := mid.Next
	mid.Next = nil

	list2 = reverseList(list2)
	head = mergeList(list1, list2)
}

// 创建链表
func buildList(nums []int) *ListNode {
	head := &ListNode{}
	p := head

	for _, num := range nums {
		p.Next = &ListNode{Val: num}
		p = p.Next
	}

	return head.Next
}

// 打印链表
func printList(list *ListNode) {
	for p := list; p != nil; p = p.Next {
		fmt.Printf("%d ", p.Val)
	}
	fmt.Println()
}

func main() {
	testCases := [][]int{
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3, 4, 5},
	}

	for _, testCase := range testCases {
		head := buildList(testCase)
		fmt.Println("原始链表：")
		printList(head)
		fmt.Println("重排链表：")
		reorderList(head)
		printList(head)
	}
}
