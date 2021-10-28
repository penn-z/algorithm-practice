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


	思路2:
		递归法
			终止条件：当两个链表都为空时，表示我们对链表已合并完成。
			如何递归：我们判断 l1 和 l2 头结点哪个更小，然后较小结点的 next 指针指向其余结点的合并结果。（调用递归）


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

// 递归合并两升序链表
func mergeTwoListsRecusion(l1 *ListNode, l2 *ListNode) *ListNode {
	// 终止条件
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		// l1头结点当前值小，则next指向 l1链表剩余节点及l2链表合并结果的头结点
		l1.Next = mergeTwoListsRecusion(l1.Next, l2)
		return l1
	} else {
		// l2头结点当前值小，则next指向 l2链表剩余节点及l1链表合并结果的头结点
		l2.Next = mergeTwoListsRecusion(l2.Next, l1)
		return l2
	}
}
