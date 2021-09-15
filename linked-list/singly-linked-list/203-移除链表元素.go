/*
	给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

	输入: head = [1,2,6,3,4,5,6], val = 6
	输出: [1,2,3,4,5]
*/
package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
   删除节点思路：
       1. 设立虚拟节点至链表最前
       2. 遍历链表，删除目标元素
*/
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head // 虚拟节点添置头结点之前
	cur := dummyHead
	// 遍历链表
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			// 链表
			cur = cur.Next
		}
	}

	return dummyHead.Next
}

func removeElementsV2(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := dummyHead
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}
