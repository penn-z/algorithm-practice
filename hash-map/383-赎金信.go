/*
	leetcode 383: 赎金信
	给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，判断第一个字符串 ransom 能不能由第二个字符串 magazines 里面的字符构成。如果可以构成，返回 true ；否则返回 false。
	(题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。杂志字符串中的每个字符只能在赎金信字符串中使用一次。)


	思路:
	1. 长度为26的数组记录ransomNote每个字符出现次数
	2. 遍历magazine，每次减去对应字符的出现次数
	3. 最后check数组每个元素的数值是否大于0，若大于领，则返回false；否则最后返回true

*/
package main

func main() {

}

func canConstruct(ransomNote string, magazine string) bool {
	record := [26]int{}

	for _, chRansom := range ransomNote {
		record[chRansom-'a']++
	}

	for _, chMagazine := range magazine {
		record[chMagazine-'a']--
	}

	for _, item := range record {
		if item > 0 {
			return false
		}
	}

	return true
}
