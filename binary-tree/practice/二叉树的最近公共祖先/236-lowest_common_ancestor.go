package main

import "fmt"

/*
	leetcode-236. 二叉树最近公共祖先
	给定一个二叉树，找到该树中两个指定节点的最近公共祖先。(假设所有节点val不重复)
	最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node2 := &TreeNode{
		Val: 2,
	}

	root := &TreeNode{
		Val:  1,
		Left: node2,
	}

	ancestor := lowestCommonAncestor(root, root, node2)
	fmt.Printf("lowest common ancestor:%v", ancestor)
}

/*
	思路:
	解法一：存储父节点
	1. 广度优先遍历二叉树，利用哈希表存储节点的父节点（只要p, q都被遍历到即可停止）
	2. 生成访问哈希表，从p节点开始不断往它的祖先移动，被访问过的节点在哈希表中标志为已访问；尔后同样从q节点开始不断往它的祖先节点移动，如果有祖先已经被访问过，即意味着这是p, q深度最深的公共祖先，LAC节点
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// BFS遍历节树
	if root == nil || p == nil || q == nil {
		return nil
	}

	parentMap := make(map[int]*TreeNode)
	visited := make(map[int]bool)
	queue := []*TreeNode{}
	queue = append(queue, root)
	for parentMap[p.Val] == nil || parentMap[q.Val] == nil {
		if len(queue) < 1 {
			break
		}

		fmt.Printf("======parentMap:%+v \n", parentMap)
		// 队头元素出队
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("========node:%d\n", node.Val)

		if node.Left != nil {
			// 加入parentMap
			parentMap[node.Left.Val] = node
			// 入队
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			// 加入paretnMap
			parentMap[node.Right.Val] = node
			// 入队
			queue = append(queue, node.Right)
		}
	}

	fmt.Printf("======final parentMap:%+v \n", parentMap)

	// p节点往上移动，往祖先节点移动
	for p != nil {
		// 标记为访问过
		visited[p.Val] = true
		// 向上父节点移动
		p = parentMap[p.Val]
	}

	// q节点向上移动，往祖先节点移动
	for q != nil {
		// 判断是否访问过
		if _, ok := visited[q.Val]; ok {
			return q
		}

		// 向上父节点移动
		q = parentMap[q.Val]
	}

	return nil
}
