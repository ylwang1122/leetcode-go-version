//n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
// 皇后会攻击同行同列同对角线上的其他皇后
//
// 
//
// 上图为 8 皇后问题的一种解法。 
//
// 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。 
//
// 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。 
//
// 示例: 
//
// 输入: 4
//输出: [
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
//解释: 4 皇后问题存在两个不同的解法。
// 
// Related Topics 回溯算法 递归求解，通过不同的尝试撤销来记录
package cn

//leetcode submit region begin(Prohibit modification and deletion)
// 皇后放置的位置 row col
var queens []int
// 当前列是否有皇后
var cols []bool
// 当前主对角线是否有皇后 同一个主对角线的节点 col - row 为常数 为避免复数因此采用 col - row + n - 1
var mains []bool
// 当前次对角线是否有皇后 同一个次对角线的节点 col + row 为常数
var secondary []bool

var output [][]string

func solveNQueens(n int) [][]string {
	if n <= 0 {
		return  nil
	}

	queens = make([]int, n)
	cols = make([]bool, n)
	mains = make([]bool, 2*n-1)
	secondary = make([]bool, 2*n-1)
	output = make([][]string, 0)

	// 利用回溯算法进行逐行尝试放置皇后
	backtrack(0, n);

	return output
}

func backtrack(row int, n int) {
	// 判定是否超限
	if row == n {
		return
	}

	// 尝试放置皇后到某一列
	for col := 0; col < n; col++ {
		// 判定是否可以防止
		if isNotUnderAttack(row, col, n) {
			// 放置皇后，更新各状态
			placeQueen(row, col, n)
			// 判定是否是最后一行，如果是则输出结果，return
			if row == n - 1 {
				// 输出行列
				addSolution(n);
				// 撤销放置皇后的操作
				removeQueen(row, col, n)
				return
			}
			// 尝试下一行操作
			backtrack(row + 1, n)
			// 撤销放置皇后的操作
			removeQueen(row, col, n)
		}
	}
}

func isNotUnderAttack(row, col int, n int) bool {
	return !cols[col] && !mains[row - col + n - 1] && !secondary[row + col]
}

func placeQueen(row, col int, n int) {
	// 添加皇后
	queens[row] = col
	// 标记列
	cols[col] = true
	// 标记主对角线
	mains[row - col + n - 1] = true
	// 标记次对角线
	secondary[row + col] = true
}

func removeQueen(row, col int, n int) {
	// 标记列
	cols[col] = false
	// 标记主对角线
	mains[row - col + n - 1] = false
	// 标记次对角线
	secondary[row + col] = false
}

func addSolution(n int) {
	// 遍历queens生成结果
	ret := make([]string, 0)
	for _, col := range queens {
		str := ""
		for i := 0; i < n; i++ {
			if i == col {
				str += "Q"
			} else {
				str += "."
			}
		}
		ret = append(ret, str)
	}
	output = append(output, ret)
}

//leetcode submit region end(Prohibit modification and deletion)
