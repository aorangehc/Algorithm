package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func result(headA, headB *ListNode) *ListNode {
	p, q := headA, headB

	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}

		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}

	return p
}

func buildList(intersectVal int, listA, listB []int, skipA, skipB int) (*ListNode, *ListNode, *ListNode) {
	var intersectNode *ListNode

	tmpA := &ListNode{}
	curA := tmpA

	for i, num := range listA {
		node := &ListNode{Val: num}
		curA.Next = node
		curA = curA.Next

		if i == skipA {
			intersectNode = node
		}
	}

	headA := tmpA.Next

	tmpB := &ListNode{}
	curB := tmpB
	for i, num := range listB {
		if i == skipB {
			break
		}
		node := &ListNode{Val: num}
		curB.Next = node
		curB = curB.Next
	}

	curB.Next = intersectNode
	headB := tmpB.Next

	return headA, headB, intersectNode
}

func main() {
	testCases := []struct {
		intersectVal int
		listA        []int
		listB        []int
		skipA        int
		skipB        int
	}{
		{
			8,
			[]int{4, 1, 8, 4, 5},
			[]int{5, 6, 1, 8, 4, 5},
			2,
			3,
		},
		{
			2,
			[]int{1, 9, 1, 2, 4},
			[]int{3, 2, 4},
			3,
			1,
		},
		{
			0,
			[]int{2, 6, 4},
			[]int{1, 5},
			3,
			2,
		},
	}

	for _, testCase := range testCases {
		headA, headB, expected := buildList(testCase.intersectVal, testCase.listA, testCase.listB, testCase.skipA, testCase.skipB)
		intersect := result(headA, headB)

		fmt.Println(expected == intersect)
	}
}
