/*
	leetcode 347:
		给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

	demo1:
		input: nums = [1,1,1,2,2,3], k = 2
		output: [1,2]
*/
package main

import "container/heap"

func main() {

}

/*
	解法一:
		求最大前k值， 利用小顶堆特性
		1. 先用map记录每个元素出现的次数
		2. 构建长度为k的小顶堆
		3. 遍历map, 若当前迭代元素的值比堆顶元素值小，则跳过，若比起大，则push入小顶堆
		4. 最后从符合题意的小顶堆中依次取出目标元素
*/
func topKFrequent(nums []int, k int) []int {
	res := []int{}
	if len(nums) == 0 || k == 0 {
		return res
	}

	numsMap := make(map[int]int)
	for _, num := range nums {
		numsMap[num]++
	}

	// 构建小顶堆
	kf := &KFrequent{}
	heap.Init(kf)

	for num, times := range numsMap {
		// 加入小顶堆
		record := &Record{val: num, times: times}
		heap.Push(kf, record)
		if kf.Len() > k {
			heap.Pop(kf)
		}
	}

	// 此时小顶堆则为符合题意的前k大元素值
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(kf).(*Record).val)
	}

	return res
}

type KFrequent struct {
	Records []*Record
}

type Record struct {
	val   int // 元素值
	times int // 元素出现次数
}

func (k *KFrequent) Less(i, j int) bool {
	return k.Records[i].times < k.Records[j].times
}

func (k *KFrequent) Swap(i, j int) {
	k.Records[i], k.Records[j] = k.Records[j], k.Records[i]
}

func (k *KFrequent) Len() int {
	return len(k.Records)
}

// pop the heap top element
func (k *KFrequent) Pop() interface{} {
	// no matter to focus on how to pop the specific element, just pop the last element, the heap package pop will deal it for us.
	old := k.Records
	val := old[len(old)-1]
	k.Records = old[:len(old)-1]
	return val
}

// push the new element into heap
func (k *KFrequent) Push(x interface{}) {
	// no need to compare the value of element, because the heap package push will deal it for us.
	k.Records = append(k.Records, x.(*Record))
}
