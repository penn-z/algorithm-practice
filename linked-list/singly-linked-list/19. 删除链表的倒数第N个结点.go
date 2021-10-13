/*
	leetcode 19: 删除链表的倒数第N个结点
	给你一个链表，删除链表的倒数第n个结点，并且返回链表的头结点。

	输入: [1, 2, 3, 4, 5], n = 2
	输出: [1, 2, 3, 5]

	思路:
		双指针，快慢指针，利用快慢指针，快指针比慢指针快n步，待快指针移动到链表末尾时，慢指针所在位置即是链表倒数第n个结点
		测试

*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

/*

 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 虚拟头结点
	dummy := &ListNode{
		Next: head,
	}

	// 快慢指针
	fast, slow := dummy, dummy

	// 快指针先行n步
	for fast != nil {
		n--
		if n >= 0 {
			fast = fast.Next
		} else {
			break
		}
	}

	// fast再行一步，需要让slow指向被删除节点的上一个节点
	fast = fast.Next

	// 快慢指针同时移动
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
