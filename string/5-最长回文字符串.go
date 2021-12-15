package main

import "fmt"

func main() {
	str := "babad"
	res := longestPalindrome(str)
	fmt.Println("====res: ", res)

}

func longestPalindrome(s string) string {
	res := ""
	// 双指针left, right，从中间向两边扩散，相等则继续
	for i := 0; i < len(s); i++ {
		// 原字符串为奇数
		// 找到以s[i]为中心的回文串
		l, r := i, i
		s1 := palindrome(s, l, r)
		if len(res) < len(s1) {
			res = s1
		}

		// 原字符串长度为偶数
		// 找到以s[i], s[i+1]为中心的回文串
		l, r = i, i+1
		s2 := palindrome(s, l, r)
		if len(res) < len(s2) {
			res = s2
		}

	}

	return res
}

func palindrome(s string, left, right int) string {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		// 向两边扩散
		left--
		right++
	}

	return s[left+1 : right]
}
