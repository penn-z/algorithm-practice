package main

func main() {

}

/*
	快慢指针
*/
func findMaxConsecutiveOnes(nums []int) int {
	fast, slow := 0, 0
	lens := len(nums)
	/*
		1. 快指针向右移，若为1，则算出当前快慢指针间隔，与已存间隔取大值
		2. 快指针指向值不为1， 则慢指针也移至快指针所在位置
	*/
	ans := 0
	for ; fast < lens; fast++ {
		if nums[fast] == 1 {
			if nums[slow] != 1 {
				slow = fast
			}

			ans = getMax((fast-slow)+1, ans)
		} else {
			slow = fast
		}
	}

	return ans
}

/*
	一次遍历:
		1. 变量a记录当前最大连续1数量，若中断则置零; 变量res记录全局最大连续数
*/

func findMaxConsecutiveOnesV2(nums []int) int {
	cnt, maxCnt := 0, 0
	for _, v := range nums {
		if v == 1 {
			cnt++
		} else {
			maxCnt = getMax(maxCnt, cnt)
			cnt = 0
		}
	}

	return maxCnt
}

func getMax(i, j int) int {
	if i < j {
		return j
	}

	return i
}
