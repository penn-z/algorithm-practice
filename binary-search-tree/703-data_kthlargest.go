package main

import (
	"container/heap"
	"sort"
)

/*
	leetcode: 703
	设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。
	请实现 KthLargest 类：
		KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
		int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。


	思路:
		小顶堆
*/

func main() {

}

type KthLargest struct {
	sort.IntSlice // sort
	k             int
}

// 构建KthLargest
func Constructtor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	// 构建堆
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (kl *KthLargest) Add(val int) int {
	// 加入堆中
	heap.Push(kl, val)

	// 若当前堆中元素数量大于题意要求k，则弹出弹出堆顶元素
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}

	// 返回intslice第一个元素，即是堆顶元素，当前堆最小元素，也是题意要求的第k大元素
	return kl.IntSlice[0]
}

// 实现heap方法 Push
func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

// Pop
func (kl *KthLargest) Pop() interface{} {
	a := kl.IntSlice
	// 弹出slice队尾元素
	v := a[len(a)-1]
	kl.IntSlice = a[:len(a)-1]
	return v
}
