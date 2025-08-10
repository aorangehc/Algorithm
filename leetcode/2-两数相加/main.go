package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	tmp := &ListNode{Val: 0}
	cur := tmp
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: carry % 10}
		carry /= 10
		cur = cur.Next
	}
	return tmp.Next

}

func printList(list *ListNode) {
	for p := list; p != nil; p = p.Next {
		fmt.Printf("%d ", p.Val)
	}
	fmt.Println()
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

func main() {
	testCases := []struct {
		nums1 []int
		nums2 []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}},
		{[]int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}},
	}

	for _, testCase := range testCases {
		list1 := buildList(testCase.nums1)
		list2 := buildList(testCase.nums2)
		listAdd := addTwoNumbers(list1, list2)
		printList(listAdd)
	}
}
