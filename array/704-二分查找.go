package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 8, 11, 15, 20}
	target := 15
	fmt.Printf("res:%v\n", binarySearch(arr, target))

	arrRightBound := []int{1, 2, 2, 2, 3, 4, 5}
	target2 := 1
	fmt.Printf("arrRightBound:%v\n", right_bound(arrRightBound, target2))
}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	return -1
}

// 搜索右边边界，左闭右开[left, right)
func right_bound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// 左闭右开区间[left, right)
	left := 0
	right := len(nums)
	// 终止条件是left == right
	for left < right {
		fmt.Println("=========")
		fmt.Printf("left: %d\n", left)
		fmt.Printf("right: %d\n", right)
		mid := (left + right) / 2
		fmt.Printf("mid: %d\n", mid)
		if nums[mid] == target {
			// 题意是要寻找最右侧边界，故不直接返回mid
			// 收缩区间左侧边界
			left = mid + 1
		} else if nums[mid] < target {
			// 目标结果在右边区间，则收缩区间左侧边界
			left = mid + 1
		} else if nums[mid] > target {
			// 目标结果在左边区间，则收缩区间右侧边界
			right = mid
		}
	}
	if left == 0 {
		return -1
	}
	if nums[left-1] == target {
		return left - 1
	} else {
		return -1
	}
}
