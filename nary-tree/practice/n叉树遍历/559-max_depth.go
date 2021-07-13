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
	return ans
}

// 迭代--DFS
/*
	1. 因为是DFS，所以显示维护一个栈来存放遍历的结点及该结点的深度
*/
func maxDepthDFS(root *Node) (ans int) {
	if root == nil {
		return
	}

	// 至少深度1
	ans = 1

	// 维护栈，存放结点
	stack := []*Node{}
	heightStack := []int{}

	stack = append(stack, root)
	heightStack = []int{1}

	// [node]depth
	nodeDepthMap := make(map[int]int)
	for len(stack) > 0 {
		// 弹出栈顶元素
		top := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		height := heightStack[len(heightStack)-1]
		heightStack = heightStack[0 : len(heightStack)-1]

		ans = max(ans, height)

		for _, child := range top.Children {
			stack = append(stack, child)
			heightStack = append(heightStack, height+1)
		}
	}

	return ans
}

// 取大值
func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
