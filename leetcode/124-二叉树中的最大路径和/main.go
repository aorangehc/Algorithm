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

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt

	var dfs func(*TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMax := dfs(node.Left)
		rightMax := dfs(node.Right)

		ans = max(ans, leftMax+rightMax+node.Val)

		return max(max(leftMax, rightMax)+node.Val, 0)
	}

	dfs(root)

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
		[]any{1, 2, 3},
		[]any{-10, 9, 20, nil, nil, 15, 7},
	}

	for _, testCase := range testCases {
		root := buildTree(testCase)
		ans := maxPathSum(root)
		fmt.Println(ans)
	}
}
