package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := []int{}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == n-1 {
				ans = append(ans, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
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
		[]any{1, 2, 3, nil, 5, nil, 4},
		[]any{1, 2, 3, 4, nil, nil, nil, 5},
		[]any{1, nil, 3},
		[]any{},
	}

	for _, testCase := range testCases {
		root := buildTree(testCase)
		ans := rightSideView(root)

		fmt.Println(ans)
	}
}
