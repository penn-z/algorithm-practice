/*
	思路：
    理解下题意，其实就是求只包含两种字符的最长子序列
    1. 使用滑动窗口windoMap；范围为[left, right]， maxCategory = 2，最长子序列长度maxLen
    2. 窗口右边界右移时，新元素加入窗口；
    3. 若当前窗口元素种类大于2，则左移左边界收缩窗口，直至窗口元素种类不大于2
    4. 记录下当前最长子序列长度maxLen
*/
package main

func main() {

}

func totalFruit(fruits []int) int {
	if len(fruits) < 1 {
		return 0
	}

	windowMap := make(map[int]int)
	left, right := 0, 0
	maxCategory := 2
	maxLen := 0

	for right < len(fruits) {
		// 右边界新元素
		rCur := fruits[right]

		// 加入窗口中
		windowMap[rCur]++

		right++
		// 是否收缩左窗口
		for len(windowMap) > maxCategory {
			// 左边界元素
			lCur := fruits[left]
			windowMap[lCur]--
			// 当前元素出现个数为0，需从窗口元素中清除
			if windowMap[lCur] == 0 {
				delete(windowMap, lCur)
			}

			left++
		}

		maxLen = getMaxVal(maxLen, right-left)
	}

	return maxLen
}

func getMaxVal(x, y int) int {
	if x >= y {
		return x
	}

	return y
}
