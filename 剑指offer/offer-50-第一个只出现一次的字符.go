/*
	offer-50:
		在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

	思路:
        1. 定义长度为26的数组，顺序记录每个字出现的次数
        2. 遍历s, 记录字母出现次数
        3. 再次遍历s，判断当前字符是否第一次出现，是则返回当前字符byte
        4. 遍历结束，无唯一字符，返回' '
*/
package main

func main() {

}

func firstUniqChar(s string) byte {
	record := make([]int, 26)
	for _, char := range s {
		record[char-'a']++
	}

	for i, char := range s {
		if record[char-'a'] == 1 {
			return s[i]
		}
	}

	return ' '
}
