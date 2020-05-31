// 给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。
// 找到并返回可以用这种方式转换的最短回文串。
//
// 示例 1: 
//
// 输入: "aacecaaa"
// 输出: "aaacecaaa"
// 
//
// 示例 2: 
//
// 输入: "abcd"
// 输出: "dcbabcd"
// Related Topics 字符串

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func shortestPalindrome(s string) string {
	// 1.暴力法，反转字符串，找到开头最长的回文子串 将后面的字符全部添加到前边即是最短的回文串
	rev := reverseStr(s)
	for i := 0; i < len(rev); i++ {
		if rev[i:] == s[:len(s) - i] {
			// 即最大的回文子串
			return rev[:i] + s
		}
	}
	return ""
	// 2. kmp算法实现
}

func reverseStr(s string) string {
	rev := ""
	for _, v := range s {
		rev = string(v) + rev
	}
	return rev
}
//leetcode submit region end(Prohibit modification and deletion)
