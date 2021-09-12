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

import (
	"fmt"
	"reflect"
)

func main() {
	// s1 := "abc"
	// s2 := "ccccbbbbaaaa"
	// s1 := "abcdxabcde"
	// s2 := "abcdeabcdx"
	// s1 := "eidboaoo"
	// s1 := "ab"
	// s2 := "eidbaooo"
	s1 := "adc"
	s2 := "dcda"
	// res := checkInclusion(s1, s2)
	// fmt.Printf("====res:%v\n", res)

	// res2 := checkInclusionV2(s1, s2)
	// fmt.Printf("====res2:%v\n", res2)

	res3 := checkInclusionV3(s1, s2)
	fmt.Printf("====res3:%v\n", res3)
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

/*
	思路2：
		1. 依然使用滑动窗口，这次固定窗口长度为len(s1)，范围是[left, right]，闭区间；needMap为s1内字符出现个数的统计，windowMap为s2内字符出现个数的统计。
		2. 右移时，需要把右边新加入窗口的字符个数在windowMap加1，把左边移出窗口的字符的个数减1。如果needMap == windowMap，则返回true退出
		3. 如果窗口已经把s2遍历完仍然没有找到，则返回false
*/
func checkInclusionV2(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	needMap, windowMap := make(map[string]int), make(map[string]int)
	for _, word := range s1 {
		needMap[string(word)]++
	}

	fmt.Printf("needMap:%+v\n", needMap)

	// 左闭右闭区间[left, right], 长度为len(s1)
	left, right := 0, len(s1)-1

	// 先把初始[left, right]放入windowMaps
	for _, item := range s2[0:right] {
		word := string(item)
		windowMap[word]++

	}

	fmt.Printf("first windowsMap:%+v\n", windowMap)

	for right < len(s2) {
		rCur := string(s2[right])
		// 把新元素放入窗口
		windowMap[rCur]++
		fmt.Printf(">>>>>add right ele, l:%d, r:%d, windowMap:%v\n", left, right, windowMap)

		if reflect.DeepEqual(windowMap, needMap) {
			// 符合题意
			return true
		}

		// 窗口向右移动前，把当前left位置的元素出现次数-1
		lCur := string(s2[left])
		windowMap[lCur]--
		fmt.Printf("<<<<<remove left ele, l:%d, r:%d, windowMap:%v\n", left, right, windowMap)
		if windowMap[lCur] == 0 {
			delete(windowMap, lCur)
		}

		// 窗口向右移动
		left++
		right++
	}

	return false
}

/*
	思路3：
		与思路2几乎一样，区别在于map的比较替换为了slice，golang底层可直接比较slice是否相等，不需要用反射

		1. 依然使用滑动窗口，这次固定窗口长度为len(s1)，范围是[left, right]，闭区间；needCounter为s1内字符出现个数的统计，windowCounter为s2内字符出现个数的统计。
		2. 右移时，需要把右边新加入窗口的字符个数在windowCounter加1，把左边移出窗口的字符的个数减1。如果needCounter == windowCounter，则返回true退出
		3. 如果窗口已经把s2遍历完仍然没有找到，则返回false


*/
func checkInclusionV3(s1, s2 string) bool {
	// base case
	if len(s1) > len(s2) {
		return false
	}

	// 初始化大小26的slice，needCounter为s1字符串元素出现个数统计， windowCounter为s2字符串元素出现个数统计， 当前元素byte相对'a'的位置则为slice的索引
	needCounter, windowCounter := [26]int{}, [26]int{}
	for _, ch := range s1 {
		needCounter[ch-'a']++
	}

	fmt.Printf("needCounter:%v\n", needCounter)

	// 左闭右闭区间[left, right], 长度为len(s1)
	left, right := 0, len(s1)-1
	// 先把初始[left, right]放入windowCounter
	for _, ch := range s2[0:right] {
		windowCounter[ch-'a']++
	}

	fmt.Printf("first windowCounter:%v\n", windowCounter)

	for right < len(s2) {
		// 把新元素放入窗口
		windowCounter[s2[right]-'a']++
		if windowCounter == needCounter {
			return true
		}

		// 窗口向右移动前，把当前left位置的元素出现次数-1
		windowCounter[s2[left]-'a']--

		// 窗口向右移动
		left++
		right++
	}

	return false
}
