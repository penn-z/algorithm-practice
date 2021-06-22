package main

import (
	"fmt"
)

// tree node struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node4 := &TreeNode{
		Val: 4,
	}

	node5 := &TreeNode{
		Val: 5,
	}

	node2 := &TreeNode{
		Val:   2,
		Right: node4,
	}

	node6 := &TreeNode{
		Val: 6,
	}

	node3 := &TreeNode{
		Val:   3,
		Left:  node5,
		Right: node6,
	}

	root := &TreeNode{
		Val:   1,
		Left:  node2,
		Right: node3,
	}

	vals := midorderTravesal(root)

	fmt.Printf("vals:%v\n", vals)

	vals = midorderIteration(root)
	fmt.Printf("vals2:%v\n", vals)

	vals = midorderIterationTemplate(root)
	fmt.Printf("vals3:%v\n", vals)
}

// 中序遍历-递归
func midorderTravesal(root *TreeNode) (vals []int) {
	var midorder func(*TreeNode)
	midorder = func(node *TreeNode) {
		// check当前节点有无val
		if node == nil {
			return
		}

		// 左子结点
		midorder(node.Left)
		// 根节点
		vals = append(vals, node.Val)
		// 右节点
		midorder(node.Right)
	}

	midorder(root)
	return
}

// 中序遍历-迭代
func midorderIteration(root *TreeNode) (vals []int) {
	// 栈记录 寻找最左节点的经过路径
	stack := []*TreeNode{}

	// 当前节点为根节点
	currNode := root
	for len(stack) > 0 || currNode != nil {
		// 找到节点最左侧，同时记录路径入栈
		for currNode != nil {
			stack = append(stack, currNode)
			currNode = currNode.Left
		}

		// 栈顶元素弹出，则为遍历的点
		top := stack[len(stack)-1]
		vals = append(vals, top.Val)
		stack = stack[:len(stack)-1]

		// 处理完最左侧节点后，判断其是否有右节点
		if top.Right != nil {
			currNode = top.Right
		}
	}

	return
}

// 僵尸模板法 //TODO: 需要理解
func midorderIterationTemplate(root *TreeNode) (vals []int) {
	// 先将根节点cur和所有左孩子入栈，直至cur为空
	stack := []*TreeNode{}

	// 根节点入栈
	curNode := root

	for curNode != nil || len(stack) > 0 {
		for curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		}

		// 每弹出一个栈顶元素，将此元素加入结果中，到达它的右孩子，并将这个节点当做cur重新执行步骤，直至栈为空
		top := stack[len(stack)-1]
		vals = append(vals, top.Val)
		stack = stack[:len(stack)-1]
		curNode = top.Right
	}

	return
}
