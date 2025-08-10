package main

import "fmt"

// 定义链表节点结构体
type ListNode struct {
	Val  int
	Next *ListNode
}

// 传入数组，创建链表
func buildList(nums []int) *ListNode {
	temp := &ListNode{} // 创建一个空节点
	p := temp

	for _, num := range nums {
		p.Next = &ListNode{Val: num}
		p = p.Next
	}

	return temp.Next
}

// 打印链表，查看效果
func printList(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Printf("%d ", p.Val)
	}
	fmt.Println()
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 统计链表长度
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n += 1
	}

	tmp := &ListNode{Next: head} // 增加哨兵，防止错误

	p := tmp
	var pre, cur *ListNode = nil, p.Next

	for n >= k {
		n -= k
		// 翻转k个节点
		for i := 0; i < k; i++ {
			nxt := cur.Next // 当前节点的后一个节点
			cur.Next = pre  // 当前节点指向前面的节点
			pre = cur       // 前一个节点指针更新
			cur = nxt       // 更新当前节点指针
		}

		// 重新将链表连接起来，链表被分成三段，已经反转好的，当前循环翻转的，后续没有翻转的
		nxt := p.Next     // 当前翻转好的最尾部的节点
		p.Next.Next = cur // 当前翻转的连接后面没有翻转的
		p.Next = pre      // 已经翻转的连接当前翻转的
		p = nxt           // p指针更新到当前翻转好的最尾部的节点
	}

	return tmp.Next
}

func main() {
	testCases := []struct {
		nums []int
		k    int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 2},
		{[]int{1, 2, 3, 4, 5}, 2},
		{[]int{1, 2, 3, 4, 5}, 3},
	}

	for _, testCase := range testCases {
		head := buildList(testCase.nums) // 创建链表
		fmt.Print("原始链表: ")
		printList(head)
		reversed := reverseKGroup(head, testCase.k)
		fmt.Print("翻转链表: ")
		printList(reversed)
	}
}
