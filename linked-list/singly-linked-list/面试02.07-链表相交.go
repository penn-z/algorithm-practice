/*
	Leetcode 面试02.07 链表相交
	给你两个单链表的头结点headA和headB，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回null

	思路:
		双指针迭代，依据题意，关键在于向右对齐两条链表尾部，从对齐点同时遍历聊表，一一对比对位节点指针是否相同即可
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 分别求出listNodeA, listNodeB链表长度
	var lenA, lenB int
	tmpA := headA
	for headA != nil {
		headA = headA.Next
		lenA++
	}
	headA = tmpA

	tmpB := headB
	for headB != nil {
		headB = headB.Next
		lenB++
	}
	headB = tmpB

	// 两者长度差绝对值
	var diff int
	if lenA > lenB {
		diff = lenA - lenB
		// headA先移动diff步
		for diff > 0 {
			headA = headA.Next
			diff--
		}
	} else {
		diff = lenB - lenA
		// headB先移动diff步
		for diff > 0 {
			headB = headB.Next
			diff--
		}
	}

	// 两条链表，分别从当前指针位置后移，一一比较两节点指针
	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}

		headA = headA.Next
		headB = headB.Next
	}

	return nil
}
