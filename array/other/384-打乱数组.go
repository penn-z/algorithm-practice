package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{1, 2, 3}
	solution := Constructor(nums)
	fmt.Println(solution.nums)
	res := solution.Shuffle()
	fmt.Println("after shuffle, res: ", res)
}

type Solution struct {
	nums   []int
	origin []int
}

func Constructor(nums []int) Solution {
	// return Solution{nums: nums, origin: nums} // 坑爹，slice是指针，底层会被改变
	return Solution{nums: nums, origin: append([]int{}, nums...)}
}

func (this *Solution) Reset() []int {
	copy(this.nums, this.origin)
	return this.nums
}

func (this *Solution) Shuffle() []int {
	// 遍历数组，(i, n-1)，每次i元素与 n-1 - i随机一元素进行交换
	n := len(this.nums)
	for i := 0; i < n; i++ {
		randomIndex := rand.Intn(n-i) + i
		this.nums[i], this.nums[randomIndex] = this.nums[randomIndex], this.nums[i]
	}

	return this.nums
}
