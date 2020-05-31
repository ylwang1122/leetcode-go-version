//给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。 
//
// 例如，给出 n = 3，生成结果为： 
//
// [
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
//]
// 
// Related Topics 字符串 回溯算法
package cn

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	// 递归实现
	// 实现递归函数 传参左右括号数量，最大数量，返回字符串
	ret := make([]string, 0)
	_generate(0, 0, n, "", &ret)
	return ret
}

// 注意传参slice的形式，直接传入可能会出现slice拷贝的问题，传入指针则没有此问题
func _generate (left, right int, n int, s string, ret *[]string) {
	// 终止条件
	if left == n && right == n {
		*ret = append(*ret, s)
		return
	}
	// 处理 过滤
	// 递归函数
	if left < n { // 左括号可以一直加到最大
		_generate(left + 1, right, n, s + "(", ret)
	}
	if right < left {
		_generate(left, right + 1, n, s + ")", ret)
	}
	// 恢复现场
}
//leetcode submit region end(Prohibit modification and deletion)
