package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var postOrder func(*TreeNode, int)
	postOrder = func(node *TreeNode, tmp int) {
		if node == nil {
			return
		}
		tmp = tmp*10 + node.Val
		if node.Left == nil && node.Right == nil {
			ans += tmp
		}
		postOrder(node.Left, tmp)
		postOrder(node.Right, tmp)

	}

	postOrder(root, 0)

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
		[]any{1, 2, 3},
		[]any{4, 9, 0, 5, 1},
	}

	for _, testCase := range testCases {
		root := buildTree(testCase)
		ans := sumNumbers(root)
		fmt.Println(ans)
	}
}
