/*
	给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。
	换句话说，s1 的排列之一是 s2 的 子串 。

	demo1:
		输入：s1 = "ab" s2 = "eidbaooo"
		输出：true
		解释：s2 包含 s1 的排列之一 ("ba").

	demo2:
		输入：s1= "ab" s2 = "eidboaoo"
		输出：false

	思路:
		1. 维护滑动窗口；needMaps记录s1目标字符串，windowMap记录当前口窗口字符及出现次数；match记录符合条件的字符个数；
		2. 右边界向右滑动:
			2.1 当前元素如果是needMaps中的目标元素，则计入windowMap且次数+1，若符合条件字符match达到len(s1)，则退出return true;否则right滑动至最右
			2.2 若当前元素否needMaps中目标元素，则清空windowMap元素及match匹配个数

*/
package main

import "fmt"

func main() {
	// s1 := "abc"
	// s2 := "ccccbbbbaaaa"
	// s1 := "abcdxabcde"
	// s2 := "abcdeabcdx"
	s1 := "ab"
	s2 := "eidboaoo"
	res := checkInclusion(s1, s2)
	fmt.Printf("====res:%v", res)
}

func checkInclusion(s1 string, s2 string) bool {
	needMap, windowMap := make(map[string]int), make(map[string]int)
	for _, word := range s1 {
		needMap[string(word)]++
	}

	// fmt.Printf("====needMap:%+v\n", needMap)

	targetLen := len(needMap)
	match := 0
	left, right := 0, 0
	s1Len, s2Len := len(s1), len(s2)
	for right < s2Len {
		rCur := string(s2[right])
		if needMap[rCur] > 0 {
			windowMap[rCur]++
			if windowMap[rCur] == needMap[rCur] {
				match++
			}

			if match == targetLen {
				// 符合题意
				return true
			}
		}

		fmt.Printf(">>>>>left:%d, right:%d, match:%d, windowMap:%v\n", left, right, match, windowMap)
		// 右边框向右移动
		right++

		// 判断左窗是否右移
		// for windowMap[rCur] > needMap[rCur] {
		for right-left >= s1Len {
			if match == targetLen {
				// 符合题意
				return true
			}

			// 提取左边框元素
			lCur := string(s2[left])

			// 左边框即将右移，若当前左边框元素为目标字符，则需要在windowMap中-1
			if needMap[lCur] > 0 {
				if windowMap[lCur] == needMap[lCur] {
					match--
				}

				windowMap[lCur]--
			}

			fmt.Printf("<<<<left:%d, right:%d, match:%d, windowMap:%v\n", left, right, match, windowMap)
			// 左边框右移
			left++
		}
	}

	return false
}
