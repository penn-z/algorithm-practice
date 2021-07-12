package main

/*
	leetcode 589: n叉树的前序遍历
*/

import "fmt"

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

	res := preorderIteration(root)
	fmt.Printf("=====res:%+v", res)
}

// 递归
func preorder(root *Node) []int {
	res := []int{}
	var preorderTraversal func(*Node, *[]int)
	preorderTraversal = func(node *Node, res *[]int) {
		// 前序遍历
		// base case
		if node == nil {
			return
		}

		// 根
		*res = append(*res, node.Val)

		// 子节点
		for _, childNode := range node.Children {
			preorderTraversal(childNode, res)
		}
	}

	preorderTraversal(root, &res)

	return res
}

// 迭代
func preorderIteration(root *Node) []int {
	// DFS
	res := []int{}

	if root == nil {
		return res
	}

	// 使用栈
	stack := []*Node{}
	stack = append(stack, root)

	// 先把根节点压入栈
	for len(stack) > 0 {
		// 栈顶节点出栈
		curNode := stack[len(stack)-1]
		if curNode != nil {
			res = append(res, curNode.Val)
		}

		// 栈顶节点出栈
		stack = stack[:len(stack)-1]

		// 因为是前序遍历，所以从右边节点开始压入栈
		if len(curNode.Children) > 0 {
			for i := len(curNode.Children) - 1; i >= 0; i-- {
				stack = append(stack, curNode.Children[i])
			}
		}

	}

	return res
}
