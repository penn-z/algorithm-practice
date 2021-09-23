package main

import "fmt"

func main() {
	/**
	 * Your MyLinkedList object will be instantiated and called as such:
	 * obj := Constructor();
	 * param_1 := obj.Get(index);
	 * obj.AddAtHead(val);
	 * obj.AddAtTail(val);
	 * obj.AddAtIndex(index,val);
	 * obj.DeleteAtIndex(index);
	 */

	obj := Constructor()
	obj.AddAtHead(1) // 1
	fmt.Printf("====head:%+v\n", obj.Get(0))

	obj.AddAtTail(3) // 1->3
	fmt.Printf("====tail:%+v\n", obj.Get(1))

	obj.AddAtIndex(1, 2) // 1->2->3
	fmt.Printf("====index:%d, val:%+v\n", 1, obj.Get(1))

	obj.DeleteAtIndex(1) // 1->3
	fmt.Printf("====index:%d, val:%+v\n", 1, obj.Get(1))
}

// 节点
type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

type MyLinkedList struct {
	head *LinkedNode
	size int // 链表长度
}

/** Initialize your val structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		head: &LinkedNode{
			Val:  0,
			Next: nil,
		},
		size: 0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}

	currentNode := this.head
	for i := 0; i <= index; i++ {
		currentNode = currentNode.Next
	}

	return currentNode.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	// 可以利用addAtIndex方法
	this.AddAtIndex(0, val)
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	// 可以利用addAtIndex
	this.AddAtIndex(this.size, val)
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	//在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。

	// index > 链表长度，不插入节点
	if index > this.size {
		return
	}

	// index < 0，在头部插入节点
	if index < 0 {
		index = 0
	}

	newNode := &LinkedNode{
		Val: val,
	}
	currentNode := this.head
	// 找到第 index 个节点的前驱节点
	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}

	newNode.Next = currentNode.Next
	currentNode.Next = newNode

	this.size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}

	currentNode := this.head
	// 找到要删除的节点
	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}

	// next跳过要删除的节点
	currentNode.Next = currentNode.Next.Next
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
