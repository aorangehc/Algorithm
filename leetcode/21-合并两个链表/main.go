package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	tmp := &ListNode{}
	p := tmp

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			p.Next = list2
			list2 = list2.Next
			p = p.Next
		} else {
			p.Next = list1
			list1 = list1.Next
			p = p.Next
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
	tmp := &ListNode{}
	p := tmp

	for _, num := range nums {
		p.Next = &ListNode{Val: num}
		p = p.Next
	}

	return tmp.Next
}

func printList(list *ListNode) {
	for p := list; p != nil; p = p.Next {
		fmt.Printf("%d ", p.Val)
	}
	fmt.Println()
}

func main() {
	testCases := []struct {
		nums1 []int
		nums2 []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}},
		{[]int{}, []int{}},
		{[]int{}, []int{0}},
	}

	for _, testCase := range testCases {
		list1 := buildList(testCase.nums1)
		list2 := buildList(testCase.nums2)
		fmt.Println("待合并链表：")
		printList(list1)
		printList(list2)

		mergedList := mergeTwoLists(list1, list2)
		fmt.Println("合并后的链表：")
		printList(mergedList)
	}
}
