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

	res := suborderTree(root)
	fmt.Println("后序遍历-递归: ", res)
	res = suborderTreeIterate(root)
	fmt.Println("后序遍历-迭代: ", res)
}

func suborderTree(root *TreeNode) (vals []int) {
	var suborder func(*TreeNode)
	suborder = func(root *TreeNode) {
		if root == nil {
			return
		}

		// 左
		suborder(root.Left)

		// 右
		suborder(root.Right)

		// 中
		vals = append(vals, root.Val)
	}

	suborder(root)
	return vals
}

// 迭代法
/*
	前序遍历顺序是 中->左->右
	后续遍历是 左->右->中
	则可以在修改前序遍历得 中->右->左，然后翻转数组即可
*/
func suborderTreeIterate(root *TreeNode) (vals []int) {
	if root == nil {
		return
	}

	// 栈记录节点
	stack := []*TreeNode{}
	// 根节点入栈
	stack = append(stack, root)
	for len(stack) > 0 {
		// 栈顶元素出栈
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 中
		vals = append(vals, top.Val)

		// 左孩子入栈
		if top.Left != nil {
			stack = append(stack, top.Left)
		}

		// 右孩子入栈
		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}

	// 翻转数组
	for i, j := 0, len(vals)-1; i < j; {
		vals[i], vals[j] = vals[j], vals[i]
		i++
		j--
	}

	return
}
