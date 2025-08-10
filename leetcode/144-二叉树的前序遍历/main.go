package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var pre func(*TreeNode)

	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		pre(node.Left)
		pre(node.Right)
	}
	pre(root)

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
		[]any{1, nil, 2, 3},
		[]any{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9},
	}

	for _, testCase := range testCases {
		root := buildTree(testCase)
		fmt.Println(preorderTraversal(root))
	}
}
