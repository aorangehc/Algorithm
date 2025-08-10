package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}

	var left, right *TreeNode

	left = lowestCommonAncestor(root.Left, p, q)
	right = lowestCommonAncestor(root.Right, p, q)

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	return root
}

// 输入节点数组，构建完整的二叉树
func buildTree(nodes []any) *TreeNode {
	// 通过层序遍历进行构建
	if len(nodes) == 0 || nodes[0] == nil { // 如果没有节点或者根节点为空，返回空树
		return nil
	}

	root := &TreeNode{Val: nodes[0].(int)} // 构建根节点，不为空，断言为int

	//var queue []*TreeNode
	//queue = append(queue, root)
	queue := []*TreeNode{root} // 先将根节点放入队列中
	i := 1
	for i < len(nodes) {
		node := queue[0]
		queue = queue[1:]
		if node == nil { // 如果是空节点跳过
			continue
		}
		if i < len(nodes) && nodes[i] != nil {
			node.Left = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, node.Left)
		}
		i += 1
		if i < len(nodes) && nodes[i] != nil {
			node.Right = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, node.Right)
		}
		i += 1
	}

	return root
}

// 查找值为val的节点
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	if left := findNode(root.Left, val); left != nil {
		return left
	}

	return findNode(root.Right, val)
}

func main() {
	testCases := []struct {
		nodes []any
		pVal  int
		qVal  int
	}{
		{[]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}, 5, 1},
		{[]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}, 5, 4},
	}

	for _, testCase := range testCases {
		root := buildTree(testCase.nodes)

		p := findNode(root, testCase.pVal)
		q := findNode(root, testCase.qVal)
		ans := lowestCommonAncestor(root, p, q)
		fmt.Println(ans)
	}
}
