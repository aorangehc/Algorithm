package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode
	ans = nil

	for _, list := range lists {
		ans = mergeTwoList(list, ans)
	}

	return ans
}

func mergeTwoList(list1, list2 *ListNode) *ListNode {
	tmp := &ListNode{}
	p := tmp

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			p.Next = list2
			p = p.Next
			list2 = list2.Next
		} else {
			p.Next = list1
			p = p.Next
			list1 = list1.Next
		}
	}

	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}

	return tmp.Next
}

func buildList(nums []int) *ListNode {
	head := &ListNode{}
	p := head

	for _, num := range nums {
		p.Next = &ListNode{Val: num}
		p = p.Next
	}

	return head.Next
}

func printList(list *ListNode) {
	for p := list; p != nil; p = p.Next {
		fmt.Printf("%d ", p.Val)
	}
	fmt.Println()
}

func main() {
	testCases := [][][]int{
		[][]int{
			[]int{1, 4, 5},
			[]int{1, 3, 4},
			[]int{2, 6},
		},
		[][]int{},
		[][]int{
			[]int{},
		},
	}

	for _, testCase := range testCases {
		testList := []*ListNode{}
		fmt.Println("原始链表：")
		for _, test := range testCase {
			list := buildList(test)
			printList(list)
			testList = append(testList, list)
		}
		mergeList := mergeKLists(testList)
		fmt.Println("合并后的链表：")
		printList(mergeList)
	}
}
