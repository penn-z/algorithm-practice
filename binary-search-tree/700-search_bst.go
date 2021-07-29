package main

/*
	给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 思路1: 递归，没啥好说的，唯一要注意题目已经给定了二叉搜索树，可降低搜索复杂地
func searchBSTRecursion(root *TreeNode, val int) *TreeNode {
	// base case
	if root == nil {
		return nil
	}

	// recusion
	// 左
	if val < root.Val {
		target := searchBSTRecursion(root.Left, val)
		if target != nil {
			return target
		}
	}

	// 中
	if val == root.Val {
		return root
	}

	// 右
	if val > root.Val {
		target := searchBSTRecursion(root.Right, val)
		if target != nil {
			return target
		}
	}

	return nil
}

// 思路2: 迭代，也没啥好说的，唯一要注意题目已经给定了二叉搜索树，可降低搜索复杂地
func searchBSTIterator(root *TreeNode, val int) *TreeNode {
	// base case
	if root == nil {
		return nil
	}

	// 维护栈
	stack := []*TreeNode{}
	curNode := root
	for curNode != nil || len(stack) > 0 {
		// 先遍历左子树
		for curNode != nil {
			if curNode.Val == val {
				return curNode
			}

			if curNode.Val > val {
				// 入栈
				stack = append(stack, curNode)
				curNode = curNode.Left
			} else {
				curNode = nil
			}

		}

		if len(stack) < 1 {
			break
		}

		// 弹栈
		top := stack[len(stack)-1]
		if top.Val == val {
			return top
		}

		if top.Val > val {
			break
		}

		// if top.Right != nil && top.Right.Val == val {
		if top.Right != nil && top.Right.Val == val {
			return top.Right
		}

		stack = stack[:len(stack)-1]
		curNode = top.Right
	}

	return nil
}
