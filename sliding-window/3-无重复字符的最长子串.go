/*
	leetcode 3: 给定一个字符串s, 请你找出其中不含有重复字符的最长子串的长度。

	demo1:
		input:   s = "abcabcbb"
		output:  3
		explain: 因为无重复字符的最长子串是"abc", 所以其长度为3

	demo2:
		input:   s = "bbbbb"
		output:  1
		explains: 无重复字符的最长子串是"b"，所以其长度为1

	思路1:
		1. 维护滑动窗口windowMap；left, right为窗口左右边界；maxLen为最长无重复子串长度
		2. 右边界右移，若新元素不在windowMap内，则加入windowMap；若新元素已存在于windowMap中，则窗口左边界左移，直至重复右边界元素出现次数为1；
		3. 移动过程不断记录窗口长度

*/
package main

import "fmt"

func main() {
	s := "pwwkew"
	res := lengthOfLongestSubstring(s)
	fmt.Printf("===res:%d\n", res)

}

/*
	思路1:
		1. 维护滑动窗口windowMap；left, right为窗口左右边界；maxLen为最长无重复子串长度
		2. 右边界右移，新元素加入windowMap内；若新元素已存在于windowMap中，则窗口左边界左移，直至重复右边界元素出现次数为1；
		3. 右边界移动过程不断记录窗口长度
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	windowMap := make(map[string]int)
	left, right := 0, 0
	maxLen := 0

	for right < len(s) {
		// 窗口右边界新元素
		rCur := string(s[right])
		windowMap[rCur]++

		// maxLen = getMaxVal(maxLen, len(windowMap))

		right++

		// 判断左窗口是否需要收缩
		for windowMap[rCur] > 1 {
			// 左边界元素
			lCur := string(s[left])
			if windowMap[lCur] > 0 {
				windowMap[lCur]--

				if windowMap[lCur] == 0 {
					delete(windowMap, lCur)
				}
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
