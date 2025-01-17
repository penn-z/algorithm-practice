package main

import "fmt"

func main() {
	str := "abcabcbb" // 3
	// str := "abba" // 2
	res := lengthOfLongestSubstring(str)
	fmt.Println("===res: ", res)

}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	// dp数组及下标含义
	// dp[i]表示以当前str[i]结尾的无重复字符子串的长度

	// 递推公式
	/*
		上一个与当前str[i]重复字符的位置即为j
		1. 如果str[i]在dp[i-1]子串中出现，则从出现位置j到下一个位置i，构成了以str[i]结尾的不重复子串:
			dp[i] = i - j
		2. 如果str[i]不在dp[i-1]子串中出现，那么dp[i] = dp[i-1] + 1
	*/

	// dp初始化
	dp := make([]int, len(s))
	dp[0] = 1

	// 哈希记录各字符最后出现的位置索引
	lastShowMap := make(map[byte]int)
	lastShowMap[s[0]] = 0

	max := 1
	for i := 1; i < len(s); i++ {
		// 获取当前字符在前缀字符串中最后出现的位置
		lastIndex, ok := lastShowMap[s[i]]
		if ok {
			if dp[i-1] < i-lastIndex {
				// 当前字符不在dp[i-1]区间内出现
				dp[i] = dp[i-1] + 1
			} else if dp[i-1] >= i-lastIndex {
				// 当前字符在dp[i-1]区间内出现
				dp[i] = i - lastIndex
			}
		} else {
			// 不在map里，故肯定不在dp[i-1]区间出现
			dp[i] = dp[i-1] + 1
		}

		lastShowMap[s[i]] = i
		if max < dp[i] {
			max = dp[i]
		}
	}

	return max
}

// tmp替换dp[i]，节省内存空间
func lengthOfLongestSubstringV2(s string) int {
	if len(s) == 0 {
		return 0
	}

	// dp数组及下标含义
	// dp[i]表示以当前str[i]结尾的无重复字符子串的长度

	// 递推公式
	/*
		上一个与当前str[i]重复字符的位置即为j
		1. 如果str[i]在dp[i-1]子串中出现，则从出现位置j到下一个位置i，构成了以str[i]结尾的不重复子串:
			dp[i] = i - j
		2. 如果str[i]不在dp[i-1]子串中出现，那么dp[i] = dp[i-1] + 1

		用map记录字符最后一次出现的位置索引


		空间优化:
			需要的是dp数组的最大值，故可用tmp存储dp[i], 变量max每轮更新最大值即可
	*/

	// dp初始化
	tmp := 0 // 替代dp[i]，节省空间
	max := 0
	lastShowMap := make(map[byte]int)
	// lastShowMap[s[0]] = 0
	for i := 0; i < len(s); i++ {
		// 获取当前字符在前缀字符串中最后出现位置
		lastIndex, ok := lastShowMap[s[i]]
		if ok {
			if tmp < i-lastIndex {
				// 当前字符不在dp[i-1]区间出现
				tmp = tmp + 1
			} else {
				// 当前字符在dp[i-1]区间出现
				tmp = i - lastIndex
			}
		} else {
			// 不在map里，故必然不在dp[i-1]区间
			tmp = tmp + 1
		}

		lastShowMap[s[i]] = i // 更新当前字符最后出现的位置索引
		if max < tmp {
			max = tmp
		}
	}

	return max
}

// 滑动窗口
/**
1. 维护滑动窗口windowMap; left, right为窗口左右边界指针；maxLen为最长无重复子串长度
2. 右边界右移，新元素加入windowMap;若新元素已存在于窗口中，则窗口左边界左移，直至右边界元素出现次数为1;
3. 右边界移动过程中不断记录窗口长度

*/
func lengthOfLongestSubstringV3(s string) int {
	if len(s) == 0 {
		return 0
	}

	windows := make(map[string]int)
	left, right, maxLen := 0, 0, 0
	for right < len(s) {
		// 元素放入窗口
		rCur := string(s[right])

		// 右移窗口
		right++

		// 进行窗口一系列数据更新
		windows[rCur]++

		// debug

		// 判断做窗口是否要收缩
		for windows[rCur] > 1 {
			// 左边界元素逻辑
			lCur := string(s[left])

			// 进行窗口内数据更新
			if windows[lCur] > 0 {
				windows[lCur]--

				if windows[lCur] == 0 {
					delete(windows, lCur)
				}
			}

			// 左移窗口
			left++
		}

		maxLen = max(maxLen, len(windows))
	}

	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
