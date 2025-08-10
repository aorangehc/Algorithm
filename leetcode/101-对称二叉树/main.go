package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var check func(*TreeNode, *TreeNode) bool
	check = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		return check(left.Left, right.Right) && check(left.Right, right.Left)
	}

	return check(root.Left, root.Right)
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
		[]any{1, 2, 2, 3, 4, 4, 3},
		[]any{1, 2, 2, nil, 3, nil, 3},
	}

	for _, testCase := range testCases {
		root := buildTree(testCase)
		fmt.Println(isSymmetric(root))
	}
}
