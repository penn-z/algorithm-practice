/*
	给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

	注意：
		对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
		如果 s 中存在这样的子串，我们保证它是唯一的答案。

	思路:
		1. 维护滑动窗口, 结果集,
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	res := minWindow(s, t)
	fmt.Printf("res:%s\n", res)
}

/*
   思路:
       1. 维护一个窗口，目标字符串map，窗口map，匹配字符串个数match，最小长度minLen
       2. 窗口右边界右移，更新数据；窗口左边界右移，更新数据。
*/

func minWindow(s string, t string) string {
	// base case
	if len(s) < 1 {
		return ""
	}

	// 窗口map, 目标字符串map
	windowsMap, needMap := make(map[string]int), make(map[string]int)
	for _, word := range t {
		needMap[string(word)]++
	}

	// left窗口左边界，right窗口右边界，match为满足needMap条件的windowsMap字符串的个数
	var left, right, match int

	// 最小子串长度
	minLen := math.MaxInt32

	// 符合题意的子串的起始索引
	start := 0

	for right < len(s) {
		// 元素移入窗口
		rEle := string(s[right])

		// 更新窗口
		if needMap[rEle] > 0 {
			windowsMap[rEle]++
			// 当前字符串满足题意出现次数
			if windowsMap[rEle] == needMap[rEle] {
				match++
			}
		}

		// 右边框右移窗
		right++

		// 判断左侧窗口是否需要收缩
		for match == len(needMap) {
			// 是否当前最小串
			if right-left < minLen {
				// 更新起始索引
				start = left
				minLen = right - left
			}

			lEle := string(s[left])

			// 左边界即将右移，若当前左边框字符为目标字符，则需要在windowsMap中减-1
			if needMap[lEle] > 0 {
				windowsMap[lEle]--
				// 若当前左边框字符次数 < 该字符目标次数，则匹配的match次数需要-1
				if windowsMap[lEle] < needMap[lEle] {
					match--
				}
			}

			// 左边框右移
			left++
		}
	}

	// 返回最小覆盖子串
	if minLen < math.MaxInt32 {
		return s[start : start+minLen]
	} else {
		return ""
	}
}
