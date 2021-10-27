/*
	leetcode 21: 合并两个有序升序链表


	思路1:
		迭代法，新定义一个新链表head，用于记录合并的结果链表
		定义cur用于指针移动。
		对l1, l2都进行遍历
			1. 当l1, l2都不为空时，比较当前l1.Val, l2.Val大小，cur.Next指向较小值，并且较小者需要指向Next进行迭代，同时cur也需要指向next迭代
			2. 当l1不为空, l2为空时，cur.Next = l1，退出迭代即可
			3. 当l1为空, l2不为空时, cur.Next = l2, 退出迭代即可
			4. 最后返回时，因为head为定义的头结点，后面才是题目要求的合并后的有序链表，故而返回head.Next
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		// l1, l2都不为nil时
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = l1
				// l1指针前进，迭代
				l1 = l1.Next
			} else {
				cur.Next = l2
				// l2指针前进，迭代
				l2 = l2.Next
			}

			// cur指针前进，迭代
			cur = cur.Next
		} else if l1 != nil {
			// l1链表不为空，l2链表为空
			cur.Next = l1 // 此时cur指向l1即可
			break
		} else {
			// l1链表为空，l2链表不为空
			cur.Next = l2 // 次数cur指向l2即可
			break
		}
	}

	// cur相当于是定义在前面的头结点，最后返回时是不需要该节点的，故而返回head.Next
	return head.Next
}
