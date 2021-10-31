/*
	leetcode 61: 旋转链表
	给你一个链表的头节点head， 旋转链表，将链表每个节点向右移动k个位置

	demo1:
		输入: head = [1, 2, 3, 4, 5], k = 2
		输出: [4, 5, 1, 2, 3]

	思路: 迭代法
		1. 首先进行一次链表遍历，找到头结点head, 尾结点tail, 链表长度len, 并求出k % len 取模结果，则是原链表头结点需要向右移动的位数
		2. 构建循环链表: cur.Next = head
		3. 再次遍历链表，找到头直接最终应该在的位置
		4. 重置头节点+断开链表
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	// 遍历链表，获取头结点，尾结点
	cur := head
	len := 1
	for cur.Next != nil {
		cur = cur.Next
		len++
	}

	k = k % len
	// 刚好循环回来，回到链表原始状态
	if k == 0 {
		return head
	}

	// 构建循环链表
	cur.Next = head

	// cur 变回头结点
	cur = head

	// 第二次遍历，找到最终头结点应该在的指针
	for i := 0; i < len-k-1; i++ {
		cur = cur.Next // 找到链表的第len-k个点
	}

	// 第len-k个点之前的是链表前n-k个节点, 之后的则是后k个节点

	// 重置头结点
	head = cur.Next

	// 断开链表，变为普通链表
	cur.Next = nil

	return head
}
