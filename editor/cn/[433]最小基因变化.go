//一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。 
//
// 假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。 
//
// 例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。 
//
// 与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。 
//
// 现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变
//化次数。如果无法实现目标变化，请返回 -1。 
//
// 注意: 
//
// 
// 起始基因序列默认是合法的，但是它并不一定会出现在基因库中。 
// 所有的目标基因序列必须是合法的。 
// 假定起始基因序列与目标基因序列是不一样的。 
// 
//
// 示例 1: 
//
// 
//start: "AACCGGTT"
//end:   "AACCGGTA"
//bank: ["AACCGGTA"]
//
//返回值: 1
// 
//
// 示例 2: 
//
// 
//start: "AACCGGTT"
//end:   "AAACGGTA"
//bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
//
//返回值: 2
// 
//
// 示例 3: 
//
// 
//start: "AAAAACCC"
//end:   "AACCCCCC"
//bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
//
//返回值: 3
// 
//

package cn

//leetcode submit region begin(Prohibit modification and deletion)
var min int
func minMutation(start string, end string, bank []string) int {
	bankMap := make(map[string]bool)
	for _, v := range bank {
		bankMap[v] = true
	}

	if !bankMap[end] {
		return -1
	}

	// dfs
	//min = -1
	//visited := make(map[string]bool)
	//dfsS(0, start, end, visited, bankMap)
	//return min

	// 双向bfs
	startMap := make(map[string]bool)
	endMap := make(map[string]bool)
	visited := make(map[string]bool)

	startMap[start] = true
	endMap[end] = true
	visited[start] = true
	visited[end] = true

	min := 0
	for len(startMap) != 0 && len(endMap) != 0 {
		if len(startMap) > len(endMap) {
			startMap, endMap = endMap, startMap
		}

		temp := make(map[string]bool)
		for k := range startMap {
			for i := 0; i < len(k); i++ {
				for _, c := range []string{"A", "C", "G", "T"} {
					s := k[0:i] + c + k[i+1:]
					if endMap[s] {
						return min + 1
					}

					if bankMap[s] && !visited[s] {
						visited[s] = true
						temp[s] = true
					}
				}
			}
		}
		startMap = temp
		min++
	}
	return -1
}

func dfsS(h int, start, end string, visited map[string]bool, bankMap map[string]bool) {
	if start == end {
		if min == -1 || min > h {
			min = h
		}
		return
	}

	visited[start] = true
	for _, v := range oneDiff(start, bankMap) {
		if visited[v] {
			continue
		}
		dfsS(h + 1, v, end, visited, bankMap)
	}
}

// 产生只变化一个字符的基因
func oneDiff(s string, bankMap map[string]bool) []string {
	ret := make([]string, 0)
	for i := range s {
		for _, b := range []string{"A", "C", "G", "T"} {
			new := s[:i] + b + s[i+1:]
			if bankMap[new] {
				ret = append(ret, new)
			}
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
