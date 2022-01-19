/*
	leetcode-1049: 最后一块石头的重量II

	有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。
	每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
		* 如果 x == y，那么两块石头都会被完全粉碎；
		* 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
	最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。

	demo1:
		输入：stones = [2,7,4,1,8,1]
		输出：1
		解释：
		组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
		组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
		组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
		组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。

	demo2:
		输入：stones = [31,26,33,21,40]
		输出：5
*/
package main

func main() {

}

func lastStoneWeight(stones []int) int {
	/*
		细看之下，该题其实是求典型的0-1背包问题。
		其实就是尽可能地把石头分成两堆，让这两堆差值尽量小，则需要两堆石头的重量接近于sum/2，最后求出的两堆石头差值即为最小值。

		则问题可以转换为： 将一堆石头放进容量为target = (sum/2)向下取整 的背包中，求放进去的石头的最大重量dp[target]，则另外一堆石头重量为sum - dp[target],
		因为target 是 sum / 2向下取整，则 sum - dp[target] <= dp[target]，
		最后所求答案即为 (sum - dp[target]) - dp[target]

		1. 背包容量为target = sum / 2，石头重量为stones[i]，价值为stones[i]
		2. dp[j]含义为 容量为j的背包，最多可以装入dp[j]重的石头
		3. 递推公式: dp[j] = max{dp[j], dp[j - stones[i]] + stones[i]}
		4. dp初始化，因为石头重量都为正整数，则dp[0...j]初始值为0
		5. 先遍历石头stones，后遍历背包，遍历背包时需要倒序
		6. 最后求(sum - dp[target]) - dp[target]，即为最后的结果
	*/

	sum := 0
	for i := range stones {
		sum += stones[i]
	}

	target := sum / 2

	dp := make([]int, target+1)

	// dp递推公式: dp = max(dp[j], dp[j - stones[i]] + stones[i])

	// dp初始化，因为石头重量为正整数，故dp[0...j]初始值为0

	// 确定遍历顺序
	for i := 0; i < len(stones); i++ {
		// 倒序遍历背包
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}

	return sum - dp[target] - dp[target]
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
