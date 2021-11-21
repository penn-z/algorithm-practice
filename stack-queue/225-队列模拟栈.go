/*
	思路1: 一个队列实现
	1. push正常入队
	2, pop时弹出队尾元素即可
*/

package main

func main() {

}

type MyStack struct {
	InQueue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.InQueue = append(this.InQueue, x)
}

func (this *MyStack) Pop() int {
	if len(this.InQueue) > 0 {
		top := this.InQueue[len(this.InQueue)-1]
		this.InQueue = this.InQueue[:len(this.InQueue)-1]
		return top
	}

	return 0
}

func (this *MyStack) Top() int {
	if len(this.InQueue) > 0 {
		return this.InQueue[len(this.InQueue)-1]
	}

	return 0
}

func (this *MyStack) Empty() bool {
	if len(this.InQueue) == 0 {
		return false
	}

	return true
}
