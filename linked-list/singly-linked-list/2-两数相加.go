/*
	leetcode 2 : 两数相加
		给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
		请你将两个数相加，并以相同形式返回一个表示和的链表。
		你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

	demo1:
		输入: l1 = [2, 4, 3], l2 = [5, 6, 4]
		输出: [7, 0, 8]
		解释: 342 + 465 = 807

	思路1: 迭代法 + 双指针法
		1. 定义个新链表头结点head, cur指向head当前节点,; carrayVal为进位值
		2. 迭代遍历l1, l2链表，对l1, l2当前指向的链表节点进行相加，得出当前位值，为cur.Val, 进位暂存，到下一位时相加
		3. 直到l1, l2指向都为空指针，停止迭代


	思路2: 递归
		终止条件: l1, l2一个为空，则返回另一个
		1. 把两个链表节点的值相加（结果记为 add ）。把 add 模 1010 作为当前的链表节点的值。
		2. 把两个链表的 next 节点相加。（注意：如果当前相加的结果 add >= 10add>=10，需要把 nextnext 节点相加得到的结果 + 1。）
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

// 迭代法
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	carryVal := 0 // 进位值
	for l1 != nil || l2 != nil {
		nowNode := &ListNode{}
		currentVal := 0
		nextCarryVal := 0
		if l1 != nil && l2 != nil {
			currentVal, nextCarryVal = getCurrentAndCarryVal(l1.Val, l2.Val, carryVal)
			carryVal = nextCarryVal
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			currentVal, nextCarryVal = getCurrentAndCarryVal(l1.Val, 0, carryVal)
			carryVal = nextCarryVal
			l1 = l1.Next
		} else {
			currentVal, nextCarryVal = getCurrentAndCarryVal(0, l2.Val, carryVal)
			carryVal = nextCarryVal
			l2 = l2.Next
		}

		nowNode.Val = currentVal
		cur.Next = nowNode

		// cur前进
		cur = cur.Next
	}

	if carryVal > 0 {
		nowNode := &ListNode{
			Val: carryVal,
		}

		cur.Next = nowNode
		cur = cur.Next
	}

	return head.Next
}

func getCurrentAndCarryVal(value1, value2, beforeVal int) (currentVal, carryVal int) {
	currentVal = value1 + value2 + beforeVal
	if currentVal >= 10 {
		currentVal = currentVal - 10
		carryVal = 1
	}

	return
}

// 两数相加-递归
/*
	终止条件: l1, l2一个为空，则返回另一个
*/
func addTwoNumbersRecursion(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	// logic

	newNode := &ListNode{}
	// 两节点值相加
	sum := l1.Val + l2.Val

	// 当前sum % 10得到的结果，即为当前node的val值
	newNode.Val = sum % 10
	// 求下一个节点
	newNode.Next = addTwoNumbersRecursion(l1.Next, l2.Next)
	// 如果当前两节点sum >= 10，则需要进位补到下一个节点值中
	if sum >= 10 {
		newNode.Next = addTwoNumbersRecursion(newNode.Next, &ListNode{Val: 1})
	}

	return newNode
}
