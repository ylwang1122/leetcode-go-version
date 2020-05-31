//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。 
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。 
//
// 
//
// 示例: 
//
// 输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
// 
//
// 说明: 
//尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。 
// Related Topics 字符串 回溯算法

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	m := map[string]string{
		"2":"abc", // 组成的字符串只能任选其一，不能一个都不选
		"3":"def",
		"4":"ghi",
		"5":"jkl",
		"6":"mno",
		"7":"pqrs",
		"8":"tuv",
		"9":"wxyz",
	}

	// 查找 digits = "234"
	// n 个格子里面只能选特定的值
	ret := make([]string, 0)
	search("", 0, digits, m, &ret)
	return ret
}

func search (str string, n int, digits string, m map[string]string, ret *[]string) {
	if n == len(digits) {
		// 追加字符串
		*ret = append(*ret, str)
		return
	}

	for _, v := range m[string(digits[n])] {
		search(str + string(v), n + 1, digits, m, ret)
	}
}
//leetcode submit region end(Prohibit modification and deletion)
