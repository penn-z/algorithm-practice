package main

/*
	leetcode 235: 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
	最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

	例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]


	输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
	输出: 6
	解释: 节点 2 和节点 8 的最近公共祖先是 6。



	思路:
        1. 利用二叉搜索树特点，左子树比当前节点小，右子树比当前节点大
        2. 利用BFS遍历，curNode节点与p, q比较，若p, q在curNode同一侧，则说明curNode为p, q祖先节点，但不是最近。若p,q不在同一侧，则curNode为最近最先节点
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}

	// BFS遍历
	queue := []*TreeNode{}

	// 根节点入队
	queue = append(queue, root)

	for len(queue) > 0 {
		// 队头元素出队
		top := queue[0]
		queue = queue[1:]

		// 同左侧，左子树入队
		if top.Val > p.Val && top.Val > q.Val {
			queue = append(queue, top.Left)
			continue
		}

		// 同右侧，右子树入队
		if top.Val < p.Val && top.Val < q.Val {
			queue = append(queue, top.Right)
			continue
		}

		// 当前节点是否p或q
		if top.Val == p.Val {
			return p
		}

		if top.Val == q.Val {
			return q
		}

		// 分叉，则当前节点为最近祖先
		return top
	}

	return nil
}
