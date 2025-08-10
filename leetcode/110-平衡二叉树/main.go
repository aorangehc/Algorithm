package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var height func(*TreeNode) int

	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(height(node.Left), height(node.Right)) + 1
	}
	return math.Abs(float64(height(root.Left)-height(root.Right))) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
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

		if i < len(nums) && nums[i] != nil {
			node.Left = &TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Left)
		}
		i += 1
		if i < len(nums) && nums[i] != nil {
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
		[]any{1, 2, 2, 3, 3, nil, nil, 4, 4},
		[]any{},
	}

	for _, testCase := range testCases {
		root := buildTree(testCase)
		fmt.Println(isBalanced(root))
	}
}
