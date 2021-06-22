package main

import "fmt"

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

	vals := subsequentorderTravesal(root)

	fmt.Printf("vals:%v\n", vals)

	vals2 := subsequentorderIterationTemplate(root)
	fmt.Printf("vals2:%v\n", vals2)
}

func subsequentorderTravesal(root *TreeNode) (vals []int) {
	var subsequentorder func(*TreeNode)

	subsequentorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 左子结点
		subsequentorder(node.Left)

		// 右子结点
		subsequentorder(node.Right)

		// 根节点
		vals = append(vals, node.Val)

	}

	subsequentorder(root)
	return
}

// 模板解法(取巧)
// 节点cur先达到最右端的叶子节点并将路径上的节点入栈
// 然后每次从栈中弹出一个元素后，达到它的左节点，并将左孩子当成cur节点，继续上述步骤
// 最后将结果反向输出即可
func subsequentorderIterationTemplate(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}

	// 根节点入栈
	curNode := root
	for curNode != nil || len(stack) > 0 {
		for curNode != nil { // 根节点、右孩子入栈
			vals = append(vals, curNode.Val)
			stack = append(stack, curNode)
			curNode = curNode.Right
		}

		// 每弹出一个栈顶元素top，就到达它的左孩子，再将这个节点当成cur重新执行上述步骤，直至栈为空
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curNode = top.Left
	}

	// 反向输出结果，双指针，一头一尾，分别向中间移动
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		vals[i], vals[j] = vals[j], vals[i]
	}

	return
}

// 模板解法(标准)
