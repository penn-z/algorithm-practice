package main

import "sort"

func arrayPairSum(nums []int) int {
	l := len(nums)
	if l < 0 {
		return 0
	}

	// 用内置sort排序即可
	sort.Ints(nums)

	total := 0
	for i := 0; i < l-1; i += 2 {
		total += nums[i]
	}

	return total
}
