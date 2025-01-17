/*
	leetcode 134: 加油站
	在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
	你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
	给定两个整数数组 gas 和 cost ，如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。


	demo1:
		输入: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
		输出: 3
		解释:
		从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
		开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
		开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
		开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
		开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
		开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
		因此，3 可为起始索引。

	demo2:
		输入: gas = [2,3,4], cost = [3,4,3]
		输出: -1
		解释:
		你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
		我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
		开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
		开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
		你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
		因此，无论怎样，你都不可能绕环路行驶一周。

	解法一：贪心策略
		总油量减去总消耗 >= 0，那么一定可以跑完一圈，说明各个站点的加油站 剩油量sum(rest[0...i])一定是 >= 0
		每个加油站剩油量rest[i] = gas[i] - cost[i]

		1. 假设经过每个加油站的油量剩余量为rest[i], 则rest[i] = gas[i] - cost[i]
		2. 定义curSum表示当前sum{rest[0...i]}，一旦curSum小于零，说明[0, i]区间都可能是起点，起始位置从i+1重新开始计算，curSum也归零
		3. 局部最优解: 当前rest[i]的和curSum一旦小于0，起始位置至少是j+1
		4. 由局部最优解推导出全局最优解： 找到可以跑一圈的起始位置
*/
package main

func main() {

}

func canCompleteCircuit(gas []int, cost []int) int {
	curSum := 0   // 当前剩余油量之和 ( >= 0)
	totalSum := 0 // 每个加油站剩余油量之和
	startIdx := 0 // 满足条件的起始位置
	for i := 0; i < len(gas)-1; i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 {
			// starIdx位置前进1
			startIdx++
			curSum = 0
		}
	}

	// 如果所有加油站剩余油量 < 0，说明汽车无法跑完一圈
	if totalSum < 0 {
		return -1
	}

	return startIdx
}
