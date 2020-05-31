//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
// 如果没有任何一种硬币组合能组成总金额，返回-1。
//
// 示例 1: 
//
// 输入: coins = [1, 2, 5], amount = 11
//输出: 3 
//解释: 11 = 5 + 5 + 1 
//
// 示例 2: 
//
// 输入: coins = [2], amount = 3
//输出: -1 
//
// 说明: 
//你可以认为每种硬币的数量是无限的。 
// Related Topics 动态规划

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	if amount < 0 || len(coins) == 0 {
		return -1
	}

	//d := make([]int, amount + 1)
	//d[0] = 0
	//for i := 1; i < amount + 1; i++ {
	//	min := amount + 1
	//	for j := 0; j < len(coins); j++ {
	//		if i >= coins[j] && d[i - coins[j]] != -1 && min > d[i - coins[j]] + 1{
	//			min = d[i - coins[j]] + 1
	//		}
	//	}
	//
	//	if min == amount + 1 {
	//		d[i] = -1
	//	} else {
	//		d[i] = min
	//	}
	//}
	// 动态规划 d[i] = min(d[i - coin[j]]) + 1
	d := make([]int, amount + 1)
	for i := 0; i < len(d); i++ {
		d[i] = amount + 1
	}
	d[0] = 0
	for i := 1; i < amount + 1; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				d[i] = minInt(d[i], d[i - coins[j]] + 1)
			}
		}
	}

	if d[amount] > amount {
		return -1
	}
	return d[amount]
}

func minInt (x, y int) int {
	if x > y {
		return y
	}
	return x
}
//leetcode submit region end(Prohibit modification and deletion)
