//给定两个单词（beginWord 和 endWord）和一个字典，找到从 beginWord 到 endWord 的最短转换序列的长度。转换需遵循如下规则：
// 
//
// 
// 每次转换只能改变一个字母。 
// 转换过程中的中间单词必须是字典中的单词。 
// 
//
// 说明: 
//
// 
// 如果不存在这样的转换序列，返回 0。 
// 所有单词具有相同的长度。 
// 所有单词只由小写字母组成。 
// 字典中不存在重复的单词。 
// 你可以假设 beginWord 和 endWord 是非空的，且二者不相同。 
// 
//
// 示例 1: 
//
// 输入:
//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出: 5
//
//解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//     返回它的长度 5。
// 
//
// 示例 2: 
//
// 输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: 0
//
//解释: endWord "cog" 不在字典中，所以无法进行转换。 
// Related Topics 广度优先搜索
package cn

//leetcode submit region begin(Prohibit modification and deletion)
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 1 广度优先搜索
	//wordMap := make(map[string]bool)
	//for _, v := range wordList {
	//	wordMap[v] = true
	//}
	//
	//if !wordMap[endWord] {
	//	return 0
	//}
	//
	//// 定义队列
	//queue := list.New()
	//visited := make(map[string]bool)
	//
	//queue.PushBack(beginWord)
	//visited[beginWord] = true
	//mix := 1
	//lastWord := beginWord
	//for (queue.Len() != 0) {
	//	fw := queue.Front()
	//	queue.Remove(fw)
	//	w := fw.Value.(string)
	//	for i := 0; i < len(w); i++ {
	//		for j := 'a'; j <= 'z'; j++ {
	//			// 使用某个字符进行替换
	//			s := w[0:i] + string(j) + w[i+1:]
	//			if s == endWord {
	//				return mix + 1
	//			}
	//			// 判定单次是否在字典中，且没被访问过
	//			if wordMap[s] && !visited[s] {
	//				visited[s] = true
	//				queue.PushBack(s)
	//			}
	//		}
	//	}
	//
	//	// 考虑是不是最后一个
	//	if lastWord == w && queue.Len() != 0{
	//		mix++
	//		lastWord = queue.Back().Value.(string)
	//	}
	//}
	//return 0

	// 2 双向bfs 从头尾进行bfs，每次只扩展数量小的一方
	wordMap := make(map[string]bool)
	for _, v := range wordList {
		wordMap[v] = true
	}

	if !wordMap[endWord] {
		return 0
	}

	// 定义队列
	beginMap := make(map[string]bool)
	endMap := make(map[string]bool)
	beginMap[beginWord] = true
	endMap[endWord] = true

	visited := make(map[string]bool)
	visited[beginWord] = true
	visited[endWord] = true

	min := 1
	for len(beginMap) != 0 && len(endMap) != 0 {
		if len(beginMap) > len(endMap) {
			beginMap, endMap = endMap, beginMap
		}

		temp := make(map[string]bool)
		for k := range beginMap {
			for i := 0; i < len(k); i++ {
				for c := 'a'; c <= 'z'; c++ {
					s := k[0:i] + string(c) + k[i+1:]
					// 判定
					if endMap[s] {
						return min + 1
					}

					if wordMap[s] && !visited[s] {
						visited[s] = true
						temp[s] = true
					}
				}
			}
		}
		beginMap = temp
		min++
	}
	return 0
}
//leetcode submit region end(Prohibit modification and deletion)
