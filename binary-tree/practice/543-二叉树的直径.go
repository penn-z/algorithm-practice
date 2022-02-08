/*
	leetcode 543 - 二叉树的直径
	给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。



示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。


	思路1： 后序遍历求所有节点的左右子树的最大深度之和

*/
package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	// 全局最大直径
	maxDiameter = 0
)

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxDiameter = 0
	getMaxDepth(root)

	return maxDiameter
}

// 计算二叉树最大深度 - 后序遍历
func getMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 求出左右子树的最大深度
	leftDepth := getMaxDepth(root.Left)
	rightDepth := getMaxDepth(root.Right)

	// 后序位置顺便求出最大直径
	sumDiameter := leftDepth + rightDepth
	maxDiameter = max(maxDiameter, sumDiameter)

	// 求出最大深度，加上根节点自身
	return max(leftDepth, rightDepth) + 1
}

func max(x, y int) int {
	return 0
}
