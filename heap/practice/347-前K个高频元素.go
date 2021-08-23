/*
	leetcode 347:
		给你一个整数组nums和一个整数k, 请你返回其中出现频率前k高的元素，你可以按任意顺序返回答案。


		解法一:
			1. 遍历一遍，使用map记录每个值出现的次数
			2. 建立大小为K的最小堆
			3. 遍历map，取出每个数字及对应出现次数，与最小堆堆顶元素对比
				1. 若次数比堆顶元素小，则跳过，继续下一个元素
				2. 若次数比堆顶元素大，则删除堆顶元素，加入当前元素
			4. 其实步骤3可变种为每次都先往堆内push元素，然后堆长度大于K时，再pop出来
			5. 最后再从符合题意的小顶堆中取出结果集输出

*/
package main

import "container/heap"

func main() {

}

type KFrequent struct {
	recordHeap []*record
}

type record struct {
	val   int
	times int
}

func topKFrequent(nums []int, k int) []int {
	if len(nums) < 1 {
		return nil
	}

	// map记录元素出现次数
	valueMap := make(map[int]int)
	for _, num := range nums {
		if _, ok := valueMap[num]; ok {
			valueMap[num]++
		} else {
			valueMap[num] = 1
		}
	}

	kf := &KFrequent{}
	heap.Init(kf)

	// 遍历map，往小顶堆写入数据
	for value, times := range valueMap {
		r := &record{val: value, times: times}
		heap.Push(kf, r)
		if kf.Len() > k {
			heap.Pop(kf)
		}
	}

	res := []int{}
	// 从小顶堆取结果集
	for i := 1; i <= k; i++ {
		res = append(res, heap.Pop(kf).(*record).val)
	}

	return res
}

func (kf KFrequent) Len() int {
	return len(kf.recordHeap)
}

func (kf KFrequent) Less(i, j int) bool {
	return kf.recordHeap[i].times < kf.recordHeap[j].times
}

func (kf KFrequent) Swap(i, j int) {
	kf.recordHeap[i], kf.recordHeap[j] = kf.recordHeap[j], kf.recordHeap[i]
}

func (kf *KFrequent) Push(val interface{}) {
	kf.recordHeap = append(kf.recordHeap, val.(*record))
	return
}

func (kf *KFrequent) Pop() interface{} {
	old := kf.recordHeap
	val := old[len(kf.recordHeap)-1]
	kf.recordHeap = old[:len(kf.recordHeap)-1]
	return val
}
