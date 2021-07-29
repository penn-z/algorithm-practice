package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}

	// 定义curNode，当前所在节点，便于插入目标节点
	curNode := root
	for {
		// 目标节小于当前节点
		if curNode.Val > val {
			// 若左子树不为空，则往左搜索
			if curNode.Left != nil {
				curNode = curNode.Left
			} else {
				// 若果左子节点为空，则直接插入，再返回root即可
				curNode.Left = &TreeNode{Val: val}
				return root
			}
		} else if curNode.Val < val {
			// 若右子树不为空，则往右搜索
			if curNode.Right != nil {
				curNode = curNode.Right
			} else {
				// 若右子节点为空，则直接插入，返回root即可
				curNode.Right = &TreeNode{Val: val}
				return root
			}

		}
	}
}
