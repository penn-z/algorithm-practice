package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 8, 11, 15, 20}
	target := 15
	fmt.Printf("res:%v\n", binarySearch(arr, target))

	arrRightBound := []int{1, 2, 2, 2, 3, 4, 5}
	target2 := 2
	fmt.Printf("arrRightBound:%v\n", right_bound_v2(arrRightBound, target2))

	left_bound_1 := []int{1, 2, 2, 2, 3, 4, 5}
	target3 := 2
	fmt.Printf("left_bound_v1:%v\n", left_bound_v1(left_bound_1, target3))

	left_bound_2 := []int{5, 7, 7, 8, 8, 10}
	target4 := 8
	fmt.Printf("left_bound_v2:%v\n", left_bound_v2(left_bound_2, target4))
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

// 搜索做左边界, 左闭右闭 [left, right]
func left_bound_v1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// 左闭右闭区间[left, right]
	left := 0
	right := len(nums) - 1
	// 终止条件是left = right + 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			// 目标结果在右区间，收缩左侧边界
			left = mid + 1
		} else if nums[mid] > target {
			// 目标结果在左区间，收缩右侧边界
			right = mid - 1
		} else if nums[mid] == target {
			// 搜索的是左侧边界，故而判断前一元素是否相等，或者当前mid为第一个元素，即为最左边界值
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				right = mid - 1
			}
		}

	}

	return -1
}

// 搜索左侧边界，左闭右开[left, right)
func left_bound_v2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// 左闭右开[left, right)
	left := 0
	right := len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			// 在右区间，收缩左侧边界
			left = mid + 1
		} else if nums[mid] > target {
			// 在左区间，收缩右侧边界
			right = mid
		} else if nums[mid] == target {
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				// 寻找做最左元素，则收缩右侧边界
				right = mid
			}
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
		mid := (left + right) / 2
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				// 若mid为最后一元素或者下一元素不等于target，则答案为当前mid
				return mid
			} else {
				// 题意是要寻找最右侧边界，故不直接返回mid
				// 收缩区间左侧边界
				left = mid + 1
			}
		} else if nums[mid] < target {
			// 目标结果在右边区间，则收缩区间左侧边界
			left = mid + 1
		} else if nums[mid] > target {
			// 目标结果在左边区间，则收缩区间右侧边界
			right = mid
		}
	}

	return -1
}

// 搜索右边边界，左闭右闭[left, right]
func right_bound_v2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// 左闭右闭区间[left, right]
	left := 0
	right := len(nums) - 1
	// 左闭右闭区间[left, right]，终止条件是left = right + 1，故而用 <=
	for left <= right {
		mid := left + (right-left)/2 // left + right / 2变体，防止int相加溢出
		if nums[mid] == target {
			if mid == len(nums)-1 || nums[mid+1] != target {
				// 若mid为最后一元素或者下一元素不等于target，则答案为当前mid
				return mid
			} else {
				// 不立即返回mid，因为要寻找右边边界，则收缩区间左侧边界，往右半区间继续搜索
				// 收缩区间左侧边界
				left = mid + 1
			}
		} else if nums[mid] < target {
			// 目标结果位于右边区间，则收缩区间左侧边界，往右半区间继续搜索
			left = mid + 1
		} else if nums[mid] > target {
			// 目标结果位于左边区间，则收缩区间右侧边界，往左半区间继续搜索
			right = mid - 1
		}
	}

	return -1
}
