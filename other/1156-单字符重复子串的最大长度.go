package main

import "fmt"

func main() {
	str := "aaabcdeafeaa"
	res := maxRepOpt1(str)
	fmt.Println("===res: ", res)
}

func maxRepOpt1(text string) int {
	// 数组存储每个元素出现次数
	record := [26]int{}
	for _, char := range text {
		record[char-'a']++
	}

	// 连续字符次数， 结果
	sequentNums, res := 1, 1

	// 上一个连续字符
	lastSequentChar := text[0]

	// 再次遍历字符串
	for i := 1; i < len(text); i++ {
		// 如果字符不连续
		if text[i] != lastSequentChar {
			// 需要检查后续字符中是否存在刚才连续字符
			rightIdx := i
			for rightIdx+1 < len(text) && lastSequentChar == text[rightIdx+1] {
				// 存在前半段的连续字符
				sequentNums++
				// 右移 右半部分游标
				rightIdx++
			}

			// 如果当前字符还未达到出现次数
			if record[lastSequentChar-'a'] > sequentNums {
				sequentNums++
			}

			res = max(res, sequentNums)
			// 重置连续字符次数
			sequentNums = 1
			lastSequentChar = text[i]
		} else {
			// 字符连续
			sequentNums++
		}
	}

	// 若sequentNums个数<该字符的总个数，sequentNums还要再+1
	if sequentNums > 1 && record[lastSequentChar-'a'] > sequentNums {
		if record[lastSequentChar-'a'] > sequentNums {
			sequentNums++
		}
	}

	return max(res, sequentNums)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
