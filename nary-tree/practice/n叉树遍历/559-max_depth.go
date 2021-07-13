package main

import "fmt"

/*
	leetcode 559 - n叉树的最大深度
*/

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	root := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val: 3,
				Children: []*Node{
					{
						Val: 5,
					},
					{
						Val: 6,
					},
				},
			},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}
	maxDepth := maxDepth(root)
	fmt.Printf("======= max depty:%d", maxDepth)
}

// 最大深度，自顶向下（前序遍历）
func maxDepth(root *Node) (ans int) {
	if root == nil {
		return
	}

	// 至少深度1
	ans = 1
	var preorder func(*Node, int)
	preorder = func(node *Node, depth int) {
		// base case
		if node == nil {
			return
		}

		// 前序遍历
		// 中
		if len(node.Children) == 0 {
			ans = max(ans, depth)
		}

		// 左, 右
		for _, childNode := range node.Children {
			preorder(childNode, depth+1)
		}
	}

	preorder(root, ans)
	return 0
}

// 最大深度，自底向上（后序遍历） TODO:
func maxDepthPoster(root *Node) (ans int) {

}

// 取大值
func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
