/*
	leetcode 206:
		给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

	输入：head = [1,2,3,4,5]
	输出：[5,4,3,2,1]

	迭代法:
		1. 定义cur记录当前节点，pre记录前一节点
		2. 迭代过程中，变量temp保存cur的下一个节点，待更新cur时使用
		3. 迭代过程中，cur.Next更新为pre; 最后移动pre与cur指针，继续下一次迭代


	递归法:
		1. 假设当前节点k后的节点都已反转，则k与k+1的反转关系为, k.next.next = k
		2.
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

// 递归法
func reverseListRecusion(head *ListNode) *ListNode {
	// base case
	if head == nil || head.Next == nil {
		return head
	}

	// 递归处理当前节点的下一个节点，得到处理好的后面节点的反转链表
	ret := reverseListRecusion(head.Next)

	// 下一个节点的next指向当前head
	head.Next.Next = head

	// 当前head.next置为nil，防止产生环
	head.Next = nil

	return ret
}
