package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	// 通过层序遍历，构建结果
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	var ans [][]int

	flag := 0
	for len(queue) > 0 {
		n := len(queue)
		var tmp []int

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]

			if node != nil {
				if flag == 0 {
					tmp = append(tmp, node.Val)
					if node.Left != nil {
						queue = append(queue, node.Left)
					}
					if node.Right != nil {
						queue = append(queue, node.Right)
					}
				} else {
					tmp = append([]int{node.Val}, tmp...)
					if node.Left != nil {
						queue = append(queue, node.Left)
					}
					if node.Right != nil {
						queue = append(queue, node.Right)
					}
				}
			}
		}
		if flag == 0 {
			flag = 1
		} else {
			flag = 0
		}
		ans = append(ans, tmp)
	}

	return ans
}

func buildTree(nums []any) *TreeNode {
	if len(nums) == 0 || nums[0] == nil {
		return nil
	}

	root := &TreeNode{Val: nums[0].(int)}
	queue := []*TreeNode{root}

	i := 1

	for i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		if nums[i] != nil && i < len(nums) {
			node.Left = &TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Left)
		}
		i += 1
		if nums[i] != nil && i < len(nums) {
			node.Right = &TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Right)
		}
		i += 1
	}

	return root
}

func main() {
	testCases := [][]any{
		[]any{3, 9, 20, nil, nil, 15, 7},
		[]any{1},
		[]any{},
	}

	for _, testCase := range testCases {
		root := buildTree(testCase)
		ans := zigzagLevelOrder(root)

		fmt.Println(ans)
	}
}
