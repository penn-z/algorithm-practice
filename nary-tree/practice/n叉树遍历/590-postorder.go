package main

import "fmt"

/*
	n叉树的后序遍历
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

	res := postorderIteration(root)
	fmt.Printf("=====res:%+v", res)
}

// 递归
func postorderRecuration(root *Node) []int {
	res := []int{}
	var postorder func(*Node, *[]int)
	postorder = func(root *Node, res *[]int) {
		// 后序遍历

		// 2 base case
		if root == nil {
			return
		}

		// 1. 框架，左右子节点
		for _, childrenNode := range root.Children {
			postorder(childrenNode, res)
		}

		// 根
		if root != nil {
			*res = append(*res, root.Val)
		}
	}

	postorder(root, &res)
	return res
}

// 迭代
/*
	与二叉树后序遍历的迭代模板法相似。
	不同之处在于:
	1. 访问二叉树左孩子节点时，N叉树改为访问第一个孩子节点
	2. 访问二叉树右孩子节点设置pre节点记录，N叉树则设置标记（即已访问过的孩子节点），然后访问下一孩子节点
	3. 判断是否访问过右孩子节点，二叉树判断pre是否为当前右孩子节点；N叉树则根据标记判断孩子节点是否全部访问过
*/
func postorderIteration(root *Node) []int {
	// DFS
	res := []int{}

	if root == nil {
		return res
	}

	// 使用栈
	stack := []*Node{}

	curNode := root
	var next int32 // 记录孩子节点访问数
	for len(stack) || curNode != nil > 0 {
		if curNode != nil {
			// 当前节点入栈
			stack = append(stack, curNode)
			curNode = curNode.Children[0]
		}

		// 栈顶节点出栈
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// stack := stack[:len(stack)-1]
		fmt.Printf("===curNode:%d\n", curNode.Val)

		if len(curNode.Children) > 0 {
			for i := len(curNode.Children) - 1; i >= 0; i-- {
				stack = append(stack, curNode.Children[i])
			}
		}

		if curNode != nil {
			res = append(res, curNode.Val)
		}
		fmt.Printf("===curNode:%+v\n", res)

	}

	return res
}

// 迭代法，后序遍历，逆序装载
func posterIterationAbnormal(root *Node) []int {
	// DFS
	res := []int{}

	if root == nil {
		return res
	}

	stack := []*Node{}

	// 根节点入栈
	stack = append(stack, root)
	for len(stack) > 0 {
		// 栈顶节点出栈
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 写入res结果数组头部(即逆序装载)
		res = append([]int{top.Val}, res...)

		// 子节点顺序压入栈
		if len(top.Children) > 0 {
			for _, childNode := range top.Children {
				stack = append(stack, childNode)
			}
		}
	}

	return res
}
