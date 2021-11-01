/*
	leetcode 242: 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
	注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

	demo1:
		输入: s = "anagram", t = "nagaram"
		输出: true

	demo2:
		输入: s = "rat", t = "car"
		输出: false

	思路: 长度为数组记录s每个字符出现次数

*/

package main

func isAnagram(s string, t string) bool {
	// 定义一个数组记录字符串s每个字符出现的次数
	record := [26]int{}

	for _, chS := range s {
		record[chS-'a']++
	}

	for _, chT := range t {
		record[chT-'a']--
	}

	for _, item := range record {
		if item != 0 {
			return false
		}
	}

	return true
}
