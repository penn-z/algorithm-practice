/*
	给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
	字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母都恰好只用一次。

	demo1:
		输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
		输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

	demo2:
		输入: strs = [""]
		输出: [[""]]

	思路:
		1. 使用长度为26的数组记录每个字符出现次数；
		2. 定义map<array, arrray>, key为上一步字符串出现次数数组，value为异位词字符串分组数组
		3. 最后遍历map，把一个key中value作为一组输出

*/
package main

func main() {

}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}

	recordMap := make(map[[26]int][]string)
	for _, str := range strs {
		// 定义长度为26的数组记录每个字符出现的次数，尔后作为key写入map中
		cnt := [26]int{}
		for _, chStr := range str {
			cnt[chStr-'a']++
		}

		recordMap[cnt] = append(recordMap[cnt], str)
	}

	res := make([][]string, 0, len(recordMap))
	for _, group := range recordMap {
		res = append(res, group)
	}

	return res
}
