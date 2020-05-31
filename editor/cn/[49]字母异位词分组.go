//给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。 
//
// 示例: 
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
//输出:
//[
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//] 
//
// 说明： 
//
// 
// 所有输入均为小写字母。 
// 不考虑答案输出的顺序。 
// 
// Related Topics 哈希表 字符串
package cn

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	// 利用map进行结果的存储 利用count进行字符的匹配验证
	//ret := make(map[string][]string)
	//
	//sort.Strings(strs)
	//for _, val := range strs {
	//	flag := false
	//	for k := range ret {
	//		if isAnagram(k, val ) {
	//			ret[k] = append(ret[k], val)
	//			flag = true
	//			break
	//		}
	//	}
	//	if !flag {
	//		ret[val] = append(ret[val], val)
	//	}
	//}
	//
	//result := make([][]string, 0)
	//for _, v := range ret {
	//	result = append(result, v)
	//}
	//return result

	m := make(map[string][]string)
	// 先排序，再判定
	for _, val := range strs {
		// 字符串排序
		s := sortString(val)
		m[s] = append(m[s], val)
	}

	rets := make([][]string, 0)
	for _, v := range m {
		rets = append(rets, v)
	}
	return rets
}

func sortString (s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] > runes[j]
	})
	return string(runes)
}

//func isAnagram(s, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//
//	counts := make([]int, 26)
//
//	for i := range s {
//		counts[s[i] - 'a']++
//		counts[t[i] - 'a']--
//	}
//
//	for _, v := range counts {
//		if v != 0 {
//			return false
//		}
//	}
//	return true
//}
//leetcode submit region end(Prohibit modification and deletion)
