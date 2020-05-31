//n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。 
//
// 
//
// 上图为 8 皇后问题的一种解法。 
//
// 给定一个整数 n，返回 n 皇后不同的解决方案的数量。 
//
// 示例: 
//
// 输入: 4
//输出: 2
//解释: 4 皇后问题存在如下两个不同的解法。
//[
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
// 
// Related Topics 回溯算法
package cn

//leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) int {
	// 1. 常规dfs 回溯
	//if n < 2 {
	//	return n
	//}
	//
	//cols := make([]bool, n) // 记录皇后放置列数
	//majority := make([]bool, 2*n - 1) // 主
	//secondary := make([]bool, 2*n - 1) // 次
	//
	//count := 0
	//dfsQueens(n, 0, &count, cols, majority, secondary)
	//return count

	// 2 结合位运算
	if n < 2 {
		return n
	}

	count := 0
	size := (1 << n) - 1 // 用于判定是否全部放置了皇后
	dfsQueens(size, 0, &count, 0, 0, 0)
	return count
}

func dfsQueens(size int, row int, count *int, cols, pie, na int) {
	if row == size { // 全部放置了皇后
		*count++
		return
	}

	// 获取剩余没有放置皇后的位置
	bits := size & (^(cols | pie | na))
	for bits != 0 {
		// 取最低为1
		p := bits & (-bits)
		bits = bits & (bits - 1)
		dfsQueens(size, row|p, count, cols|p, (pie | p) << 1, (na | p) >> 1)
	}
}

//func dfsQueens(n, row int, count *int, cols, majority, secondary []bool) {
//	if row == n {
//		*count++
//		return
//	}
//
//	for col := 0; col < n; col++ {
//		// 判定当前列是否可以放置
//		if !cols[col] && !majority[n - col + row - 1] && !secondary[col + row] {
//			// 放置皇后
//			cols[col] = true
//			majority[n - col + row - 1] = true
//			secondary[col + row] = true
//			dfsQueens(n, row + 1, count, cols, majority, secondary)
//			// 撤销皇后
//			cols[col] = false
//			majority[n - col + row - 1] = false
//			secondary[col + row] = false
//		}
//	}
//}
//leetcode submit region end(Prohibit modification and deletion)
