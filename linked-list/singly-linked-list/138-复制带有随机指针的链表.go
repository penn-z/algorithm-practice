/*
	leetcode: 138 - 复制带有随机指针的链表
	给你一个长度为 n 的链表，每个节点包含一个额外增加的随机指针 random ，该指针可以指向链表中的任何节点或空节点。
	构造这个链表的 深拷贝。 深拷贝应该正好由 n 个 全新 节点组成，其中每个新节点的值都设为其对应的原节点的值。新节点的 next 指针和 random 指针也都应指向复制链表中的新节点，并使原链表和复制链表中的这些指针能够表示相同的链表状态。复制链表中的指针都不应指向原链表中的节点 。

	思路: 迭代法 + map记录旧新指针映射关系
	1. 迭代旧链表过程中，用map[oldNode]newNode 记录映射关系
	2. 迭代完成后，再次遍历新链表，填充Random指针即可
*/

package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func main() {

}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// oldNode => newNode
	oldNewMap := make(map[*Node]*Node)

	cur := head
	newHead := &Node{Val: cur.Val}
	oldNewMap[cur] = newHead
	newCur := newHead
	for cur.Next != nil {
		newNode := &Node{
			Val: cur.Next.Val,
		}

		// 写入map
		oldNewMap[cur.Next] = newNode

		// 赋值Next
		newCur.Next = newNode

		// 指针移动
		newCur = newCur.Next
		cur = cur.Next
	}

	// cur, newCur指针已到尾部，重新回到头部
	cur = head
	newCur = newHead
	// 填充新链表每个结点的random指针
	for cur != nil {
		newCur.Random = oldNewMap[cur.Random]
		cur = cur.Next
		newCur = newCur.Next
	}

	return newHead
}
