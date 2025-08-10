package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNthNode(head *ListNode, n int) *ListNode {
	tmp := &ListNode{Next: head}

	slow, fast := tmp, tmp

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	
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
	testCases := []struct {
		nums []int
		n    int
	}{
		{[]int{1, 2, 3, 4, 5}, 2},
		{[]int{1}, 1},
		{[]int{1, 2}, 1},
	}

	for _, testCase := range testCases {
		head := buildList(testCase.nums)
		fmt.Println("原始链表：")
		printList(head)
		deletedList := deleteNthNode(head, testCase.n)
		fmt.Println("删除链表：")
		printList(deletedList)
		fmt.Println("----------------")
	}
}
