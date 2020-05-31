//实现 pow(x, n) ，即计算 x 的 n 次幂函数。 
//
// 示例 1: 
//
// 输入: 2.00000, 10
//输出: 1024.00000
// 
//
// 示例 2: 
//
// 输入: 2.10000, 3
//输出: 9.26100
// 
//
// 示例 3: 
//
// 输入: 2.00000, -2
//输出: 0.25000
//解释: 2-2 = 1/22 = 1/4 = 0.25 
//
// 说明: 
//
// 
// -100.0 < x < 100.0 
// n 是 32 位有符号整数，其数值范围是 [−2^31, 2^31 − 1] 。
// 
// Related Topics 数学 二分查找

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func myPow(x float64, n int) float64 {
	/*
		pow(x^n)
			subproblem: subresult pow(x^n/2)
		merge:
			if n & 1 == 1 then subresult = subresult * subresult * x (奇数)
			else subresult = subresult * subresult
	*/
	// 处理复数的情况
	if n < 0 {
		n = -n
		return 1/fastPow(x, n)
	}
	return fastPow(x, n)
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	sub := fastPow(x, n/2)
	if n & 1 == 1 {
		sub = sub * sub * x
	} else {
		sub = sub * sub
	}
	return sub
}
//leetcode submit region end(Prohibit modification and deletion)
