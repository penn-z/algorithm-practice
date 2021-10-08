/*
	leetcode 206:
		给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

	输入：head = [1,2,3,4,5]
	输出：[5,4,3,2,1]
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseList(head *ListNode) *ListNode {
	// 定义cur记录当前节点
	cur := head
	// 定义pre记录前一节点
	var pre *ListNode
	for cur != nil {
		// 定义temp保存cur的下一个节点
		temp := cur.Next
		cur.Next = pre
		// 更新pre指针和cur指针
		pre = cur
		cur = temp
	}

	return pre
}
