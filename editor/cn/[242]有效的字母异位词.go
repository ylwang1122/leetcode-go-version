//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。 
//
// 示例 1: 
//
// 输入: s = "anagram", t = "nagaram"
//输出: true
// 
//
// 示例 2: 
//
// 输入: s = "rat", t = "car"
//输出: false 
//
// 说明: 
//你可以假设字符串只包含小写字母。 
//
// 进阶: 
//如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？ 
// Related Topics 排序 哈希表
package cn

//leetcode submit region begin(Prohibit modification and deletion)
// 1 问题讨论 什么是异位词？
// 2 解决方案 --> 最优解决方案
// 3 code
// 4 测试用例
func isAnagram(s string, t string) bool {
	//
	//if len(s) != len(t) {
	//	return false
	//}
	//
	//m := make(map[byte]int)
	//for i := range []byte(s) {
	//	m[s[i]]++
	//	m[t[i]]--
	//}
	//
	//for _, v := range m {
	//	if v != 0 {
	//		return false
	//	}
	//}
	//
	//return true

	if len(s) != len(t) {
		return false
	}

	// 使用26个字符表示的数组进行统计 前提都是小写
	counts := make([]int, 26)
	for i := range t {
		counts[s[i] - 'a']++
		counts[t[i] - 'a']--
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
