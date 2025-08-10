package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func trainingPlan(head *ListNode, cnt int) *ListNode {
	slow, fast := head, head

	for i := 0; i < cnt; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
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
		nums []int
		cnt  int
	}{
		{[]int{2, 4, 7, 8}, 1},
	}

	for _, testCase := range testCases {
		head := buildList(testCase.nums)
		ans := trainingPlan(head, testCase.cnt)
		fmt.Println(ans.Val)
	}
}
