package main

import "fmt"

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

	res := preorderTree(root)
	fmt.Println("前序遍历-递归: ", res)

	res = preorderTreeIterate(root)
	fmt.Println("前序遍历-迭代: ", res)

}

func preorderTree(root *TreeNode) (vals []int) {
	if root == nil {
		return
	}

	var preorder func(*TreeNode)

	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		// 中
		vals = append(vals, root.Val)
		// 左
		preorder(root.Left)
		// 右
		preorder(root.Right)
	}

	preorder(root)
	return
}

// 迭代法
/*
	维护一个栈，先加入右孩子节点，再加入左孩子节点，最后加入根节点。
	则出栈时的顺序为 中 -> 左 -> 右
*/
func preorderTreeIterate(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}

	// 根节点入栈
	stack = append(stack, root)
	for len(stack) > 0 {
		// 栈顶元素出栈
		top := stack[len(stack)-1]
		if top != nil {
			vals = append(vals, top.Val)
		}

		stack = stack[:len(stack)-1]

		// 右孩子节点入栈
		if top.Right != nil {
			stack = append(stack, top.Right)
		}

		// 左孩子节点入栈
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}

	return
}
