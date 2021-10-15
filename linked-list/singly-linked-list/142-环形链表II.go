/*
	leetcode 142: 环形链表II
	给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
	为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
	说明：不允许修改给定的链表。

	思路1:
		遍历链表，使用map记录每个结点指针地址，若重复出现，则返回该节点
		空间复杂度: O(n)
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func detectCycle(head *ListNode) *ListNode {
	// 定义map，记录结点指针地址
	pointerMap := make(map[*ListNode]bool)

	// 虚拟节点
	dummy := &ListNode{
		Next: head,
	}

	cur := dummy
	for cur.Next != nil {
		if _, ok := pointerMap[cur.Next]; ok {
			return cur.Next
		}

		pointerMap[cur.Next] = true
		cur = cur.Next
	}

	return nil
}
