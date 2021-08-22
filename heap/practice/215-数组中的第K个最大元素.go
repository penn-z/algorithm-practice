/*
	leetcode-215: 数组中的第K个最大元素

	思路:
		1. 使用最小堆即可求解
		2. 创建大小为K的最小堆
		3. 依次将元素添加到最小堆中;
		4. 当最小堆的元素到达K个时，将当前元素与堆顶元素对比:
			1. 若当前元素小于堆顶元素，则放弃当前元素，继续进行下一个元素
			2. 若当前元素大于堆顶元素，则删除堆顶元素，将当前元素加入到最小堆中
		5. 重复3，4步骤，直到所有元素遍历完毕
		6. 此时，最小堆中的堆顶元素就是第K个最大的元素。
*/
package main

import (
	"container/heap"
	"sort"
)

type KthLarest struct {
	sort.IntSlice // 默认是升序int数组
	k             int
}

func main() {

}

func findKthLargest(nums []int, k int) int {
	kth := &KthLarest{k: k}
	var res int
	for _, num := range nums {
		res = kth.Add(num)
	}

	return res
}

func (kth *KthLarest) Add(val int) int {
	// 加入堆中
	heap.Push(kth, val)

	// 若当前堆中元素个数大于K，则需要弹出堆顶元素
	if kth.Len() > kth.k {
		heap.Pop(kth)
	}

	// 返回kth.IntSlice第一个元素，堆顶元素，即是第K大元素

	return kth.IntSlice[0]
}

func (kth *KthLarest) Pop() interface{} {
	a := kth.IntSlice
	// 弹出队尾元素
	v := a[len(a)-1]
	kth.IntSlice = a[:len(a)-1]
	return v
}

func (kth *KthLarest) Push(value interface{}) {
	kth.IntSlice = append(kth.IntSlice, value.(int))
	return
}
