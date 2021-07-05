package main

type TreeTreeNode struct {
	Val   int
	Left  *TreeTreeNode
	Right *TreeTreeNode
	Next  *TreeTreeNode
}

func main() {

}

// 不使用BFS，空间复杂度O(1)
func connect(root *TreeTreeNode) *TreeTreeNode {
	if root == nil {
		return root
	}

	//cur我们可以把它看做是每一层的链表
	cur := root

	// 外层循环，层级切换
	for cur != nil {
		// 遍历当前层时，为方便切换至下一层，可在下一层前面添加哑节点dummy
		dummy := &TreeNode{}

		// pre表示为下一层节点的前一位置
		pre := dummy
		// 内层循环，在当前层遍历
		for cur != nil {
			// 判断当前节点left是否为空，不为空则TreeNode就是下一层第一个节点了
			if cur.Left != nil {
				pre.Next = cur.Left
				pre = pre.Next
			}

			// 判断当前节点right是否为空，不为空则TreeNode就是下一层第一个节点了
			if cur.Right != nil {
				pre.Next = cur.Right
				pre = pre.Next
			}

			// 当前节点右移
			cur = cur.Next
		}

		// 当前节点下移
		cur = dummy.Next
	}

	return root
}
