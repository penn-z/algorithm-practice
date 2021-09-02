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

func getMax(i, j int) int {
	if i < j {
		return j
	}

	return i
}
