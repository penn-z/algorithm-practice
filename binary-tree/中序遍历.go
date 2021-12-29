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

	res := midorderTree(root)
	fmt.Println("递归中序遍历: ", res)

	res = midorderTreeIterate(root)
	fmt.Println("中序遍历-递归: ", res)
}

func midorderTree(root *TreeNode) (vals []int) {
	var midorder func(*TreeNode)
	midorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		// 左
		midorder(root.Left)

		// 中
		vals = append(vals, root.Val)

		// 右
		midorder(root.Right)
	}

	midorder(root)
	return vals
}

// 迭代法
/*
	1. 定义指针访问树节点，栈则处理节点上的元素
*/
func midorderTreeIterate(root *TreeNode) (vals []int) {
	if root == nil {
		return
	}

	stack := []*TreeNode{}
	// 访问树节点的指针
	cur := root
	for cur != nil || len(stack) > 0 {
		// 指针访问到最底层
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left // 左
		}

		// 栈顶元素出栈
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 加入到结果集中
		vals = append(vals, top.Val) // 中

		// 处理完了最左节点，指针遍历右结点
		cur = top.Right // 右

	}

	return
}
