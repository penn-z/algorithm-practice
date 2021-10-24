/*
	leetcode 328: 奇偶链表
	给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。请注意，这里的奇数节点和偶数节点指的是节点编号和奇偶性，而不是节点的值的奇偶性。
	请尝试使用原地算法完成。你的算法的空间复杂度为O(1)，时间复杂度为O(nodes), nodes为节点总数。


	思路:
		1. 使用奇偶双指针，当遍历完链表后，偶数链表
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 因为都是使用指针, 遍历odd过程，相当于head为odd的头结点
	// even头结点需要使用另外变量even_head记录
	odd := head
	even := head.Next
	even_head := even // 保存当前偶数结点的头结点
	for odd.Next != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next

		odd = odd.Next
		even = even.Next
	}

	odd.Next = even_head // 拼接链表
	// 最后返回head即可，因为head为odd头结点
	return head
}
