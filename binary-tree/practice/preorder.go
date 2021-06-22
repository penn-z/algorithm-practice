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

	// vals := preorderTraversal(root)
	vals := preorderTraversal(root)

	fmt.Printf("traversal vals:%v\n", vals)

	vals2 := preorderIteration(root)
	fmt.Printf("iteration vals:%v\n", vals2)

	vals3 := preorderIterationTemplate(root)
	fmt.Printf("iteration val3:%v\n", vals3)

}

// 递归前序遍历
func preorderTraversal(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}

	preorder(root)
	return
}

// 迭代前序遍历
// 维护一个栈，模拟递归过程，前序遍历是根结点 -> 左子结点 -> 右子结点
// 故而当前结点出栈后，先压右结点入栈，后压左结点入栈
func preorderIteration(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}

	// 先把根结点压入栈
	stack = append(stack, root)

	for len(stack) > 0 {
		// 栈顶节点
		curNode := stack[len(stack)-1]
		if curNode != nil {
			vals = append(vals, curNode.Val)
		}

		// 栈顶节点出栈
		stack = stack[:len(stack)-1]

		// 右节点入栈
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}

		// 左节点入栈
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
	}

	return
}

// 僵尸模板解法
func preorderIterationTemplate(root *TreeNode) (vals []int) {
	// 先将根节点cur和所有左孩子入栈并加入结果中，直至cur为空，有个一个while循环实现
	stack := []*TreeNode{}

	curNode := root
	// stack = append(stack, root)

	for curNode != nil || len(stack) > 0 {
		for curNode != nil { // 根节点、左孩子入栈
			vals = append(vals, curNode.Val)
			stack = append(stack, curNode)
			curNode = curNode.Left
		}

		// 每弹出一个栈顶元素top，就到达它的右孩子，再将这个节点当做cur重新按照上面的步骤来一遍，直至栈为空
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curNode = top.Right
	}

	return
}
