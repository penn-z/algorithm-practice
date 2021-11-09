/*
	leetcode 454 四数相加 II
	给你四个整数数组 nums1, num2, num3, num4, 数组长度都是n, 请你计算有多少个元组(i, j, k, l)满足:
		* 0 <= i, j, k, l < n
		* nums1[i] + num2[j] + num3[k] + num4[l] == 0

	示例:
		输入: num1 = [1, 2], num2 = [-2, -1], num3 = [-1, 2], num4 = [0 ,2]
		输出: 2
		解释: 两个元组如下:
		1. (0, 0, 0, 1) => nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
		2. (1, 1, 0, 0) => nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0

	思路: 利用hash即可. map<plus, times>, key 为 两数组之和，times为该和出现次数
		1. 遍历nums1, nums2,两两组合计算 a + b之和，记录入map<plus, times>中
		2. 遍历num3, nums4, 两两组合计算0 - (c+d)值作为key是否出现在map<plus, times>中，出现则取得value值，累加值res结果
		3. 最后返回结果res即可


*/
package main

func main() {

}

func fourSumCouont(nums1, nums2, nums3, nums4 []int) int {
	var res int
	plusMap := make(map[int]int)
	// 遍历nums1, nums2
	for _, a := range nums1 {
		for _, b := range nums2 {
			plusMap[a+b]++
		}
	}

	// 遍历nums3, nums4, 计算 0 - (c +d)出现在plusMap中次数，累加至结果中
	for _, c := range nums3 {
		for _, d := range nums4 {
			if count, has := plusMap[0-(c+d)]; has {
				res = res + count
			}
		}
	}

	return res
}
