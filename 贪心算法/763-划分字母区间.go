/*
	leetcode 763
	字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。

	demo:
		输入：S = "ababcbacadefegdehijhklij"
		输出：[9,7,8]
		解释：
		划分结果为 "ababcbaca", "defegde", "hijhklij"。
		每个字母最多出现在一个片段中。
		像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。


	思路:
	1. 遍历字符串，记录下每个字符的最大索引maxIndex
	2. 再次遍历字符串，判断已扫描过的字符串，每个字符最大maxIndex 是否出现在当前分区右边界索引之前：
        是则继续遍历
        否则需要分段
*/

package main

import "fmt"

func main() {
	s := "ababcbacadefegdehijhklij"
	res := partitionLabels(s)
	fmt.Println("res: ", res)
}

func partitionLabels(s string) []int {
	res := []int{}
	if len(s) == 0 {
		return res
	}

	// 1. 遍历字符串，记录每个字符出现的最大索引下标
	maxIndexDict := [26]int{}
	for i, char := range s {
		maxIndexDict[char-'a'] = i
	}

	/*
		2. 再次遍历字符串，check已扫描区间，每个字符的最大maxIndex 是否出现在当前分区的右边界索引之前：
			是则继续遍历
			否则需要更新当前分区右边界下标值
	*/
	start := 0                // 分区左边界
	curPartitionMaxIndex := 0 // 分区右边界
	for i, char := range s {
		curMaxIndex := maxIndexDict[char-'a']
		// 当前字符出现的最大索引 大于 当前分区右边界，需要更新分区右边界
		if curMaxIndex > curPartitionMaxIndex {
			curPartitionMaxIndex = curMaxIndex
		}

		// 扫描下标i 与 分区右边界相遇，则需要执行分区
		if i == curPartitionMaxIndex {
			res = append(res, i-start+1)
			// 更新分区左边界
			start = i + 1
		}
	}

	return res
}
