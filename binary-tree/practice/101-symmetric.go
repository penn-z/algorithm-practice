// 对称二叉树

package main

import "fmt"

// tree node struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	root := buildTree()
	isSymmetric := isSymmetricTree(root)
	fmt.Printf("recursion isSymmetric:%v\n", isSymmetric)

	isSymmetric = isSymmetricTreeIteration(root)
	fmt.Printf("iteration isSymmetric:%v\n", isSymmetric)
}

func buildTree() (root *TreeNode) {
	node4 := &TreeNode{
		Val: 3,
	}

	node5 := &TreeNode{
		Val: 4,
	}

	node2 := &TreeNode{
		Val:   2,
		Left:  node4,
		Right: node5,
	}

	node6 := &TreeNode{
		Val: 4,
	}

	node7 := &TreeNode{
		Val: 3,
	}

	node3 := &TreeNode{
		Val:   2,
		Left:  node6,
		Right: node7,
	}

	root = &TreeNode{
		Val:   1,
		Left:  node2,
		Right: node3,
	}

	return
}

// 是否对称二叉树 (递归)
func isSymmetricTree(root *TreeNode) (isSymmetric bool) {
	// 递归终止条件:
	// 1. 两个节点都为空，则当前子树为镜像对称
	// 2. 两个节点一个为空，当前子树不为镜像对称
	// 3. 当前两节点值不相等
	// 3. 递归比较左右子树，left.Left != right.Right 或 left.Right != right.Left, 则不为镜像对称
	var symmetircTree func(left *TreeNode, right *TreeNode) bool
	symmetircTree = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			// 递归是求解子节点的答案，故子节点对称
			return true
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		return symmetircTree(left.Left, right.Right) && symmetircTree(left.Right, right.Left)
	}

	isSymmetric = symmetircTree(root.Left, root.Right)
	return
}

// 是否镜像对称二叉树（迭代方法）
func isSymmetricTreeIteration(root *TreeNode) (isSymmetric bool) {
	// 当连个子树根节点相等时，就比较
	// 左子树的left和右子树的right，这里可以用队列来实现
	// 1. 从队列中拿出两个节点(left 和 right)比较
	// 2. 将left.Left和right.Right节点放入队列
	// 3. 讲left.Right和right.Left节点放入队列

	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	// 生成队列保存节点
	queue := []*TreeNode{}
	// 将根节点左右孩子加入队列
	queue = append(queue, root.Left, root.Right)
	// queue = append(queue, root, root)

	// 遍历队列
	for len(queue) > 0 {
		// 从队头取出两个节点
		l, r := queue[0], queue[1]
		queue = queue[2:]

		// 两节点是否同时为nil，若真，则表示该两点均为叶子节点的子节点，跳过，继续循环遍历其他节点即可
		if l == nil && r == nil {
			// 迭代法使用的是BFS，此时不一定遍历完所有节点，故而继续循环
			continue
		}

		// 某一节点为空，不对称
		if l == nil || r == nil {
			return false
		}

		// 对比节点值不相等，不对称
		if l.Val != r.Val {
			return false
		}

		// 此时两对比节点对称，则继续比较各自子节点
		queue = append(queue, l.Left, r.Right)
		queue = append(queue, l.Right, r.Left)
	}

	return true
}
