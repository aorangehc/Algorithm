package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		n := len(queue)
		tmp := make([]int, len(queue))

		for i := 0; i < n; i++ {
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}

			tmp[i] = queue[0].Val
			queue = queue[1:]
		}

		ans = append(ans, tmp)
	}

	return ans
}

// 构建树
func buildTree(nodes []any) *TreeNode {
	if len(nodes) == 0 || nodes[0] == nil {
		return nil
	}

	root := &TreeNode{Val: nodes[0].(int)}
	queue := []*TreeNode{root} // 先将根节点放进队列
	i := 1

	for i < len(nodes) {
		cur := queue[0]
		queue = queue[1:]

		if i < len(nodes) && nodes[i] != nil {
			cur.Left = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, cur.Left)
		}
		i += 1

		if i < len(nodes) && nodes[i] != nil {
			cur.Right = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, cur.Right)
		}
		i += 1
	}

	return root
}

func main() {
	testCases := [][]any{
		[]any{1, 9, 20, nil, nil, 15, 7},
		[]any{1},
		[]any{},
		[]any{1, 2, 3, 4, 5, nil, 7},
	}

	for _, testCase := range testCases {
		fmt.Println("测试用例")
		fmt.Println(testCase)

		root := buildTree(testCase)
		result := levelOrder(root)

		fmt.Println("遍历结果")
		fmt.Println(result)
	}
}
