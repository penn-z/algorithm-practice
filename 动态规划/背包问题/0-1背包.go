/*
	0-1背包问题：
	假设有重量为[1, 3, 4]的三个物品，价值为[15, 20, 30]，背包重量为4，求背包能装下的最大价值是多少。

	思路: 动态规划
		1. 确定dp数组及其下标含义，dp[i][j]表示从[0-i]的物品里任意取，放进容量为j的背包中，价值总和的最大值.
		2. 确定递推公式，dp[i][j]可由
			dp[i-1][j], 不取物品i，因为取了物品i后，重量总和会大于j；
			dp[i-1][j - weight[i]] + value[i]，取物品i，取了物品i后，总重量不会超过j
		3. dp初始化
			dp[i][0]，表示背包容量j=0，则什么也装不下，dp[i][0] = 0
			dp[0][j]表示，装第一个物品，
				若j < weight[0]，第一个物品装不下，dp[0][j] = 0;
				若j >= weight[0]，区间为[weight[0], begWeight]，则dp[0][j] = value[0]

		4. 确定遍历顺序
			先遍历物品，后遍历背包
			for i := 1; i < len(weight); i++ { // 遍历物品
				for j := 0; j <= weight; j++ {

				}
			}


*/
package main

import "fmt"

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	// res := package0_1(weight, value, 4)
	res := package0_1_compress(weight, value, 4)
	fmt.Println("res: ", res)
}

func package0_1(weight, value []int, bagWeight int) int {
	if len(weight) == 0 {
		return 0
	}

	if bagWeight == 0 {
		return 0
	}

	/*
		1. 确定dp数组及下标含义：
			dp[i][j]表示从[0, j]的物品任意取，放进容量为j的背包中，价值总和的最大值.
	*/

	/*
		2. 递推dp[i][j]公式，
			dp[i][j] = max{
				dp[i-1][j], // 表示无法放入物品i，会超过背包容量j，此时背包中的总价值
				dp[i-1][j - weight[i]] + value[i] // 表示可以放入物品i，此时背包红的总价值
			}
	*/
	/*
		3. dp初始化
			当背包容量为0时，dp[i][0] = 0

			当取第一个物品时，dp[0][j]
				若 j < weight[0]，第一个物品无法放入
					dp[0][j] = 0
				若 j >= weight[0], 取值区间为[weight[0], begWegiht],
					dp[0][j] = value[0]
	*/

	/*
		4. 确定遍历顺序
			先遍历物品，后遍历背包
			for i := 1; i < len(weight); i++ {
				for j := 0; j <= begWeight; j++ {

				}
			}
	*/

	dp := make([][]int, len(weight))

	for i := range dp {
		dp[i] = make([]int, bagWeight+1)
	}

	// dp初始化，当背包容量为0时，背包价值为0
	for i := 0; i < len(weight); i++ {
		dp[i][0] = 0
	}

	for j := 0; j <= bagWeight; j++ {
		if j < weight[0] {
			// 背包容量小于第一个物品weight
			dp[0][j] = 0
		} else {
			// 背包容量不小于第一个物品weight
			dp[0][j] = value[0]
		}
	}

	// 确定遍历顺序
	// 先遍历物品，后遍历背包
	for i := 1; i < len(weight); i++ {
		// 遍历背包
		for j := weight[i]; j <= bagWeight; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
		}
	}

	return dp[len(weight)-1][bagWeight]
}

// 0_1背包状态压缩，二维数组压缩为一维数组
func package0_1_compress(weight, value []int, bagWeight int) int {
	if len(weight) == 0 {
		return 0
	}

	if bagWeight == 0 {
		return 0
	}

	/*
		对于二维数组dp[i][j]，dp[i][j] = max{dp[i-1][j], dp[i-1][j - weight[i]] + value[i]}
		发现如果把dp[i-1]这一层拷贝到dp[i]这一层上，dp[i-1]可以不再使用，则 dp[i][j] = max{dp[i][j], dp[i][j - weight[i]] + value[i]}，
		发现其实i这一维完全可以不用使用，直接压缩为dp[j]一维滚动数组即可。

		这个滚动数组，满足的条件是上一层可以重复利用，直接拷贝到当前层！！！

		1. 则在dp[j]中，含义为容量为j的背包，所背的物品总价值最大值为dp[j]。

		2. 递推公式
			dp[j] = max{
				d[j], // 不放物品i
				dp[j - weight[i]] + value[i],  // 放置物品i
			}

		3. dp初始化
			由dp[j] = max{dp[j], dp[j - weight[i]] + value[i]} 可得知dp推导过程中一定是取价值最大的数，故dp[j]所有下标对应的dp值都为0即可

		4. 确定遍历顺序
			需要从后往前倒序遍历，是为了保证物品i只被放入一次！如果一旦正序遍历了，那么物品0就会被重复加入多次！

			为什么二维数组dp[i][j]不需要倒序遍历呢？因为对于二维数组dp，dp[i][j]都是通过上一层dp[i-1][j]计算得来，本层的dp[i][j]并不会被覆盖！
	*/

	dp := make([]int, bagWeight+1)

	// 递推公式 dp[j] = max(dp[j], dp[j - weight[i]] + value[i])

	// dp各元素值初始化为0即可

	for i := 0; i < len(weight); i++ {
		// 倒序遍历dp[j]，防止覆盖
		for j := bagWeight; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}

	fmt.Println("dp: ", dp)

	return dp[bagWeight]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
