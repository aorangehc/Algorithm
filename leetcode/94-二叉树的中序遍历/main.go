// leetcode-94-二叉树的中序遍历
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var inorder func(node *TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)

		ans = append(ans, node.Val)

		inorder(node.Right)
	}

	inorder(root)

	return ans
}

func preTraversal(root *TreeNode) []int {
	ans := []int{}

	var pre func(node *TreeNode)
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

func postTraversal(root *TreeNode) []int {
	ans := []int{}

	var post func(node *TreeNode)
	post = func(node *TreeNode) {
		if node == nil {
			return
		}
		post(node.Left)
		post(node.Right)
		ans = append(ans, node.Val)
	}

	post(root)
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
		[]any{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9},
		[]any{1, nil, 2, 3},
		[]any{},
		[]any{1},
	}

	for _, testCase := range testCases {
		root := buildTree(testCase)
		ansPre := preTraversal(root)
		ansIn := inorderTraversal(root)
		ansPost := postTraversal(root)
		fmt.Println(ansPre, ansIn, ansPost)
	}
}
