//给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。如果剩余少于 k 个字符，则将剩余的所有全部反转。如果有小于
// 2k 但大于或等于 k 个字符，则反转前 k 个字符，并将剩余的字符保持原样。 
//
// 示例: 
//
// 
//输入: s = "abcdefg", k = 2
//输出: "bacdfeg"
// 
//
// 要求: 
//
// 
// 该字符串只包含小写的英文字母。 
// 给定字符串的长度和 k 在[1, 10000]范围内。 
// 
// Related Topics 字符串
package cn

//leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	// i,i+k    i+2k,i+3k  i+4k,i+5k
	start, end := 0, k - 1
	for end < len(s) {
		// 反转字符串
		reverse(&s, start, end)
		start = end + k + 1
		end = start + k - 1
	}

	if start < len(s) {
		if end < len(s) {
			reverse(&s, start, end)
		} else {
			reverse(&s, start, len(s) - 1)
		}
	}

	return s
}

func reverse(s *string, start, end int) {
	runes := []rune(*s)
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
	*s = string(runes)
}
//leetcode submit region end(Prohibit modification and deletion)
