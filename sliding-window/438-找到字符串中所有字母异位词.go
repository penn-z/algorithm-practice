package main

func main() {

}

/*
   思路1:
       1. 利用固定大小的滑动窗口winLen := len(p), 范围是[left, right]闭区间; needCounter为p内字符出现个数的统计; windowCounter为s内字符出现个数的统计；
       2. 右移时，需要把右边新加入窗口的字符在windowCounter中+1，把左边移除窗口的字符在windCounter中-1
       3. 比较needCounter == windowCounter是否成立，成立则计入结果集中，否则不计入
*/
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	// 初始化结果集
	ans := []int{}

	// 分别初始化大小为26的slice: needCounter, windowCounter
	needCounter, windowCounter := [26]int{}, [26]int{}
	for _, ch := range p {
		needCounter[ch-'a']++
	}

	// 初始窗口索引位置
	left, right := 0, len(p)-1

	// 初始化窗口
	// 这里初始化的窗口长度少1
	for _, ch := range s[left:right] {
		windowCounter[ch-'a']++
	}

	for right < len(s) {
		// 右边界值加入窗口
		windowCounter[s[right]-'a']++
		// 若此时windowCounter与needCounter相等，则加入结果集
		if windowCounter == needCounter {
			ans = append(ans, left)
		}

		// 窗口向右移动前，左边界元素移除
		windowCounter[s[left]-'a']--

		// 窗口向右移动
		left++
		right++
	}

	return ans
}
