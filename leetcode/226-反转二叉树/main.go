package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	var invert func(*TreeNode)

	invert = func(node *TreeNode) {
		if node == nil {
			return
		}
		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp

		invert(node.Left)
		invert(node.Right)
	}

	invert(root)

	return root
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

// 层序遍历打印
func printTree(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i] == nil {
				continue
			}
			fmt.Printf("%d ", queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
	}
	fmt.Println()
}

func main() {
	testCases := [][]any{
		[]any{4, 2, 7, 1, 3, 6, 9},
		[]any{2, 1, 3},
		[]any{},
	}

	for _, testCase := range testCases {
		root := buildTree(testCase)
		printTree(invertTree(root))
	}
}
