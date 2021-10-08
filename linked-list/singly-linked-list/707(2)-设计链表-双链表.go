package main

import "fmt"

func main() {

	obj := Constructor()
	obj.AddAtHead(1) // 1
	// fmt.Println("size: ", obj.size)
	fmt.Printf("====head:%+v\n", obj.Get(0))

	// obj.AddAtTail(3) // 1->3
	// fmt.Printf("====tail:%+v\n", obj.Get(1))

	// obj.AddAtIndex(1, 2) // 1->2->3
	// fmt.Printf("====index:%d, val:%+v\n", 1, obj.Get(1))

}

type LinkedNode struct {
	Val  int
	Prev *LinkedNode
	Next *LinkedNode
}

type MyLinkedList struct {
	head *LinkedNode // 哨兵头节点
	tail *LinkedNode // 哨兵尾结点
	size int
}

func Constructor() MyLinkedList {
	head := &LinkedNode{}

	tail := &LinkedNode{}

	head.Next = tail
	tail.Prev = head

	return MyLinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	// 通过比较index和size - index的大小判断从头开始还是从尾开始比较快
	// 从较快的方向开始

	if index < 0 || index >= this.size {
		return -1
	}

	currentNode := this.head
	if index+1 < this.size-index {
		// 从头开始比较快
		for i := 0; i <= index; i++ {
			currentNode = currentNode.Next
		}
	} else {
		// 从尾开始比较快
		currentNode = this.tail
		for i := 0; i <= this.size-index; i++ {
			currentNode = currentNode.Prev
		}
	}

	return currentNode.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	// 可以直接利用addAtIndex方法
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	// 可以直接利用addAdIndex方法
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// 在链表中的第 index 个节点之前添加值为 val 的节点。
	// 若 index等于链表长度，则该节点附加到链表的末尾；
	// 若 index小于0，则在头部插入节点

	if index > this.size {
		return
	}

	if index < 0 {
		index = 0
	}

	newNode := &LinkedNode{
		Val: val,
	}

	currentNode := this.head
	if index+1 < this.size-index {
		// 从头开始查找比较快
		for i := 0; i <= index; i++ {
			currentNode = currentNode.Next
		}
	} else {
		// 从尾开始查找比较快
		currentNode = this.tail
		for i := 0; i <= this.size-index; i++ {
			currentNode = currentNode.Prev
		}
	}

	newNode.Next = currentNode.Next
	newNode.Prev = currentNode
	currentNode.Next.Prev = newNode
	currentNode.Next = newNode
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	// 越界
	if index < 0 || index >= this.size {
		return
	}

	currentNode := this.head
	if index+1 < this.size-index {
		// 从头开始查找比较快
		for i := 0; i <= index; i++ {
			currentNode = currentNode.Next
		}
	} else {
		// 从尾开始查找比较快
		currentNode = this.tail
		for i := 0; i <= this.size-index; i++ {
			currentNode = currentNode.Prev
		}
	}

	// 跳过要删除的节点
	currentNode.Next = currentNode.Next.Next
	currentNode.Next.Prev = currentNode
	this.size--
}
