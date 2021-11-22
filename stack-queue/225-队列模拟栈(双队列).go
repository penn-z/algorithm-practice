/*
	思路: 两个队列实现: 队列q1操作实际出队元素，队列q2仅用作备份
*/

package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	param_3 := obj.Top()
	fmt.Println("param_3:", param_3)
	param_2 := obj.Pop()
	fmt.Println("param_2:", param_2)
	param_4 := obj.Empty()
	fmt.Println("param_4:", param_4)
}

type MyStack struct {
	InQueue     []int
	BackupQueue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.InQueue = append(this.InQueue, x)
}

func (this *MyStack) Pop() int {
	if len(this.InQueue) > 0 {
		this.in2Backup()
		res := this.InQueue[0]
		this.InQueue = []int{}
		this.backup2In()
		return res
	}

	return 0
}

func (this *MyStack) Top() int {
	if len(this.InQueue) > 0 {
		this.in2Backup()
		res := this.InQueue[0]
		this.backup2In()
		return res
	}

	return 0
}

func (this *MyStack) Empty() bool {
	if len(this.InQueue) == 0 {
		return true
	}

	return false
}

func (this *MyStack) in2Backup() {
	if len(this.InQueue) > 1 {
		this.BackupQueue = append(this.BackupQueue, this.InQueue[0:len(this.InQueue)-1]...)
		// 留最后一个
		this.InQueue = []int{this.InQueue[len(this.InQueue)-1]}
	}
}

func (this *MyStack) backup2In() {
	if len(this.BackupQueue) > 0 {
		this.InQueue = append(this.BackupQueue, this.InQueue...)
		this.BackupQueue = []int{}
	}
}
