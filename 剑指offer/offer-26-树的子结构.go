/**
leetcode: offer-26-树的子结构

输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
B是A的子结构， 即 A中有出现和B相同的结构和节点值。

demo1:
	input: A = [1, 2, 3], B = [3, 1]
	output: false

demo2:
	input: A = [3, 4, 5, 1, 2], B = [4, 1]
	output: true

思路:
	若树B是树A的子结构，则子结构的根节点可能是树A的任意节点，故判断B是否A的子结构需要判断如下条件:
	1. A的根节点是否包含B
		a. 节点B为空，越过了B，为A子结构
		b. 节点B为空，节点A不为空，越过了A，则B不为A子结构
		c. A,B均不为空，但A,B节点值不相等，则B不为A子结构
		d. A,B均不为空且当前节点相等，则继续递归判断各自左右节点
	2. A的左右节点为根节点的子树是否包含B
		a. 递归调用A左子树、右子树


	可知，上述1 -> 2为前序遍历


*/

package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	var isSame func(*TreeNode, *TreeNode) bool
	isSame = func(nodeA *TreeNode, nodeB *TreeNode) bool {
		// 返回条件
		// nodeB为空，越过了B，可知B为A子结构
		if nodeB == nil {
			return true
		}

		// B不为空，A却是空，此时越过了A，B不为A子结构
		if nodeA == nil {
			return false
		}

		// A, B都不为空，但是值不相等，B不为A子结构
		if nodeA.Val != nodeB.Val {
			// A, B节点值不相等，返回false
			return false
		}

		// A左右子树是否与B左右子树相同
		return isSame(nodeA.Left, nodeB.Left) && isSame(nodeA.Right, nodeB.Right)
	}

	// import! 这里是前序遍历
	// 1. 先判断根节点是否有B子结构

	// A为根节点结构是否与B相同
	if isSame(A, B) {
		return true
	}

	// 2. 再判断A左子树，及A右子树是否有b结构
	// A左子树是否与B一致，B右子树是否与B一致
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
