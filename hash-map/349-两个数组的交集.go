/*
	leetcode 349: 两个数组的交集
	给定两个数组，编写一个函数计算它们的交集。

	demo1:
		输入: nums1 = [1,2,2,1], nums2 = [2,2]
		输出: [2]

	demo2:
		输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
		输出: [9,4]

	思路:
		1. 利用set特性，定义set1, set2，分别记录两个数组的值
		2. set1, set2中，长度较短的set作为外循环，check另一个set2内是否存在对应元素，有则加入结果数组中
*/
package main

func main() {

}

func interestion(nums1 []int, nums2 []int) []int {
	set1 := map[int]struct{}{}
	for _, num := range nums1 {
		set1[num] = struct{}{}
	}

	set2 := map[int]struct{}{}
	for _, num := range nums2 {
		set2[num] = struct{}{}
	}

	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}

	res := []int{}
	for num, _ := range set1 {
		if _, has := set2[num]; has {
			res = append(res, num)
		}
	}

	return res
}
