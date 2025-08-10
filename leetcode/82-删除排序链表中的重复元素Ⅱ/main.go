package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// tmp := &ListNode{Next: head}
	pre, cur := head, head.Next

	for cur != nil {
		if pre.Val == cur.Val {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			pre = pre.Next
			cur = cur.Next
		}
	}

	return head
}

func deleteAllList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tmp := &ListNode{Next: head}

	// 删除所有重复的
	pre, cur := tmp, head
	for cur != nil && cur.Next != nil {
		val := cur.Val
		if val == cur.Next.Val {
			for cur != nil && cur.Val == val {
				cur = cur.Next
			}
			// 不重复，跳出循环
			pre.Next = cur
		} else {
			pre = pre.Next
			cur = cur.Next
		}

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
		[]int{1, 2, 3, 3, 4, 4, 5},
		[]int{1, 1, 1, 2, 3},
	}

	for _, testCase := range testCases {
		head := buildList(testCase)
		fmt.Println("原始链表：")
		printList(head)
		head1 := head
		fmt.Println("处理后链表1：")
		deletedList := deleteList(head)
		printList(deletedList)
		fmt.Println("处理后链表2：")
		deletedAllList := deleteAllList(head1)
		printList(deletedAllList)
		fmt.Println("-------------------")

	}
}
