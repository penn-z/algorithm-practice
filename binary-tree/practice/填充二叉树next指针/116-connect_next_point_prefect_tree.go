package main

/*
	leetcode-116: 给定一个 完美二叉树(满二叉树) ，其所有叶子节点都在同一层，每个父节点都有两个子节点。
	填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
	初始状态下，所有 next 指针都被设置为 NULL。

	思路:
		1. 可用BFS实现
		2. 当前
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

func main() {

}

func connect(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	// 初始化队列
	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		// 每一层数量
		levelCount := len(queue)

		// 前一节点
		preNode := &TreeNode{}
		for i := 0; i < levelCount; i++ {
			node := queue[0]
			queue = queue[1:]

			// 若前一节点为空，则当前节点为该层第一个，不处理
			if preNode != nil {
				preNode.Next = node
			}

			// 当前节点成为前一节点
			preNode = node

			// 左子树入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			// 右子树入队
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}

// 不使用BFS，空间复杂度为O(1)
func connectV2(root *TreeNode) *TreeNode {
	/*
		思路:
			完美二叉树即满二叉树
			1. 遍历当前层时，可把下一层通过链表链接
			2. 两种连接类型：
				1. 同一父节点下的左右孩子节点，node.Left.next = node.Right
				2. 不同父节点下的右左孩子节点，及左父节点右孩子连接右父节点左孩子，可通过当前层的左父节点next找到右父节点，即node.Right = node.next.Left
			3. 遍历完该第N层后，改变层数，下移一层，继续遍历，为第N层建立next指针时，处于第N-1层。当第N层节点的next指针全部建立后，移至第N层，建立第N+1层的next指针。

		算法:
	*/

	if root == nil {
		return root
	}

	// 每层遍历从最左节点开始
	var leftMostNode *TreeNode = root

	for leftMostNode.Left != nil {

		node := leftMostNode
		for node != nil {
			// 左子树节点next指向右子树节点
			node.Left.Next = node.Right

			// 右子树节点next指向当前节点next节点的左子树节点
			node.Right.Next = node.Next.Left

			// 当前节点右移
			node = node.Next
		}

		// 当前层数下移，切换层数
		leftMostNode = leftMostNode.Left
	}

	return root
}
