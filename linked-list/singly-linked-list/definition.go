package main

type Node struct {
	Data int
	Next *Node
}

type List struct {
	HeadNode *Node // 头结点
}
