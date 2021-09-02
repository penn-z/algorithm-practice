package main

// removeElement
func removeElement(nums []int, val int) int {
	fast, slow := 0, 0
	lens := len(nums)
	for ; fast < lens; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
