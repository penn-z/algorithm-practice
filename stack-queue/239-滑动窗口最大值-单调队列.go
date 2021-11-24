/*
	思路1： 暴力法
		1. 维护一个定长窗口，窗口移动时，每次求得窗口内的最大值, 复杂度为O(n * k), k为窗口移动次数

	思路2: 最大堆---不得行
		1. 建立一个数量为k的最大堆，每次移动窗口，则把新元素放入堆中，堆顶即为最大值
		2. 但是移动窗口时，左边界的值无法从最大堆中删除，故此法不通

	思路3: 单调队列，设计单调队列结构，包含push(), pop(), front()方法
		1. 维护一个单调递增的队列，队头出口元素即为将弹出的元素，为队列中最大值
		2. 移动窗口时，
			弹出窗口的元素调用单调队列pop():
				如果窗口移除的元素value == 单调队列的出口元素value，则将该出口元素弹出，否则不做操作
		   	加入窗口的元素调用单调队列push(value):
			   	如果push的元素值大于入口元素值，则将队列入口的元素弹出，直到push元素值小于等于单调队列入口元素值为止
		3. 保持上述规则移动，每次移动时，调用front()即可获取当前窗口的最大值。
*/
package main

func main() {

}

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}

	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	// 滑动窗口 + 单调队列
	// 创建一个单调队列
	queue := Constructor()
	// 前k个元素用来填充队列
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}

	// 前k个元素的最大值加入结果集中
	res = append(res, queue.Front())

	// 遍历nums剩余元素
	for i := k; i < len(nums); i++ {
		// 滑动窗口移除最左边元素
		queue.Pop(nums[i-k])
		// 滑动窗口加入最右边元素
		queue.Push(nums[i])
		// 记录当前窗口中的最大值
		res = append(res, queue.Front())
	}

	return res
}

// 单调队列
type MonotoneQueue struct {
	queue []int
}

func Constructor() MonotoneQueue {
	return MonotoneQueue{}
}

func (m *MonotoneQueue) Push(num int) {
	// 如果窗口元素新入的元素value 大于 队列的入口元素，则弹出入口元素。直到窗口新入元素value <= 队列入口元素
	for len(m.queue) > 0 && num > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}

	// 窗口新元素加入单调队列中
	m.queue = append(m.queue, num)

}

func (m *MonotoneQueue) Pop(num int) {
	// 如果窗口移除的元素value == 单调队列的出口元素value，则将该出口元素弹出，否则不做操作
	if len(m.queue) > 0 && m.Front() == num {
		m.queue = m.queue[1:]
	}
}

// 当前队列的出口元素，即为最大值
func (m *MonotoneQueue) Front() int {
	return m.queue[0]
}

// 当前队列的入口元素
func (m *MonotoneQueue) Back() int {
	return m.queue[len(m.queue)-1]
}
