package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BFS(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	q := []*TreeNode{root}

	// 下面的循环为模版代码
	for len(q) != 0 {
		n := len(q)
		for _, node := range q {
			res = append(res, node.Val) // 替换为实际代码

			if node.Left != nil { // 注意判断非空
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[n:]
	}
	return res
}
