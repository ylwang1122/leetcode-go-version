//给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换
//需遵循如下规则： 
//
// 
// 每次转换只能改变一个字母。 
// 转换过程中的中间单词必须是字典中的单词。 
// 
//
// 说明: 
//
// 
// 如果不存在这样的转换序列，返回一个空列表。 
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
//输出:
//[
//  ["hit","hot","dot","dog","cog"],
//  ["hit","hot","lot","log","cog"]
//]
// 
//
// 示例 2: 
//
// 输入:
//beginWord = "hit"
//endWord = "cog"
//wordList = ["hot","dot","dog","lot","log"]
//
//输出: []
//
//解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。 
// Related Topics 广度优先搜索 数组 字符串 回溯算法
package cn

import "container/list"

//leetcode submit region begin(Prohibit modification and deletion)
type treeNode struct {
	word string
	count int
	pre *treeNode
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordMap := make(map[string]bool)
	for _, v := range wordList {
		wordMap[v] = true
	}

	if wordMap[endWord] == false {
		return nil
	}

	rets := make([][]string, 0)

	queue := list.New()
	visited := make(map[string]bool)
	queue.PushBack(&treeNode{
		word: endWord,
		count: 0,
		pre:  nil,
	})
	flag := false // 记录是否查询成功
	count := 0
	for queue.Len() != 0 {
		f := queue.Front()
		queue.Remove(f)
		w := f.Value.(*treeNode)

		if flag && count != w.count {
			continue
		}

		// 变形单词
		visited[w.word] = true
		for i := 0; i < len(w.word); i++ {
			for j := 'a'; j <= 'z'; j++ {
				s := w.word[0:i] + string(j) + w.word[i+1:]
				if s == beginWord {
					// 逆序查找
					strs := make([]string, 0)
					strs = append(strs, beginWord)
					cur := w
					for cur.pre	!= nil {
						strs = append(strs, cur.word)
						cur = cur.pre
					}
					strs = append(strs, cur.word)
					rets = append(rets, strs) //  只记录当前层的数据
					flag = true
					count = w.count
				}

				if wordMap[s] && !visited[s] {
					queue.PushBack(&treeNode{
						word: s,
						count:w.count + 1,
						pre:  w,
					})
				}
			}
		}
	}
	return rets
}
//leetcode submit region end(Prohibit modification and deletion)
