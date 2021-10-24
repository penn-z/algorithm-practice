/*
	leetcode: 876. 链表的中间结点
	给定一个头结点为head的非空单链表，返回链表的中间结点。
	如果有两个中间结点，则返回第二个中间结点。

	demo1:
		输入: [1, 2, 3, 4, 5]
		输出: 此列表中的结点3 (序列化形式: [3, 4, 5])，返回的结点值为3。

	思路:
		利用快慢指针特性
		1. 当快指针到达末端时，此时slow即是中点所在位置， 返回slow即可。
		2. 注意: 判断循环条件需要看fast.Next是否为空，因为题意要求最后中点有两个时，需要返回第二个。
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
