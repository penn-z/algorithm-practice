/*
	leetcode 24: 两两交换链表中的节点
	给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
	你不能只是单纯地改变节点内部的值，而是需要实际地进行节点交换。

	输入： head = [1, 2, 3, 4]
	输出： [2, 1, 4, 3]

	务必要画图，执行步骤更清晰
	迭代法:
		使用虚拟头结点dummy处理
		1. 两两节点进行处理：节点a, b； cur表示当前处理节点，初始指向dummy
		2. 假设原本节点为 dummy -> 1 -> 2 -> 3 -> 4
			步骤一: dummy -> 2
			步骤二: dummy -> 2 -> 1
			步骤三: dummy -> 2 -> 1 -> 3
		3. cur右移两个节点，继续下一轮迭代
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	// 定义虚拟头节点dummy
	dummy := &ListNode{
		Next: head,
	}

	cur := dummy
	// 迭代
	for cur.Next != nil && cur.Next.Next != nil {
		tmp1 := cur.Next
		tmp2 := cur.Next.Next.Next

		cur.Next = cur.Next.Next  // 步骤一
		cur.Next.Next = tmp1      // 步骤三: 前面需要变量记录
		cur.Next.Next.Next = tmp2 // 步骤二: 前面需要变量记录

		// cur右移两位
		cur = cur.Next.Next
	}

	return dummy.Next
}
