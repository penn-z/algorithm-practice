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

	vals3 := subsequentorderIterationTemplateV2(root)
	fmt.Printf("vals3:%v\n", vals3)
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

	// 根节点开始
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
func subsequentorderIterationTemplateV2(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	pre := &TreeNode{} // 用于记录前一个访问的节点

	// 根节点开始
	curNode := root
	for curNode != nil || len(stack) > 0 {
		for curNode != nil {
			// 当前节点入栈
			stack = append(stack, curNode)
			curNode = curNode.Left // 左
		}

		// 弹出一个栈顶元素top，就到达它的右孩子，再将这个节点当成cur重新执行上述步骤，直至栈为空
		// 如何访问根节点top呢？？如果右子树为空或者已经访问过了，则此时访问根节点

		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if top.Right != nil && top.Right != pre {
			// 再次压回栈中
			stack = append(stack, top)
			// 访问右子树
			curNode = top.Right // 右
		} else {
			// 访问根节点
			vals = append(vals, top.Val) // 根
			pre = top                    // 更新访问记录
			// top = nil
		}
	}

	return vals
}
