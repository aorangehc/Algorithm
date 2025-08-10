package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var maxCount func(*TreeNode) int
	maxCount = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max(maxCount(node.Left)+1, maxCount(node.Right)+1)
	}

	return maxCount(root)
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
		[]any{1, nil, 2},
	}

	for _, testCase := range testCases {
		root := buildTree(testCase)
		fmt.Println(maxDepth(root))
	}
}
