package main

/*
	给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。

	demo1:
		输入: nums = [1,2,3,1], k = 3
		输出: true

	demo2:
		输入: nums = [1,0,1,1], k = 1
		输出: true

	demo3:
		输入: nums = [1,2,3,1,2,3], k = 2
		输出: false


	思路:
		1. 维护map[int]int, 记录元素值与下标映射
		2. 遍历nums，若当前num不在map中，加入map
		3. 若当前num在map中，判断当前下标与map中记录下标绝对值，是否<=k，是则返回true退出；否则更新当前map中对应值下标映射
*/

// containsNearbyDuplicate
func containsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := indexMap[num]; ok {
			if abs(i, j) <= k {
				return true
			} else {
				indexMap[num] = i
			}
		} else {
			indexMap[num] = i
		}
	}

	return false
}

func abs(x, y int) int {
	r := x - y
	if r < 0 {
		return -r
	}

	return r
}
