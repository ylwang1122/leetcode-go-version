//班上有 N 名学生。其中有些人是朋友，有些则不是。他们的友谊具有是传递性。如果已知 A 是 B 的朋友，B 是 C 的朋友，那么我们可以认为 A 也是 C 
//的朋友。所谓的朋友圈，是指所有朋友的集合。 
//
// 给定一个 N * N 的矩阵 M，表示班级中学生之间的朋友关系。如果M[i][j] = 1，表示已知第 i 个和 j 个学生互为朋友关系，否则为不知道。你
//必须输出所有学生中的已知的朋友圈总数。 
//
// 示例 1: 
//
// 
//输入: 
//[[1,1,0],
// [1,1,0],
// [0,0,1]]
//输出: 2 
//说明：已知学生0和学生1互为朋友，他们在一个朋友圈。
//第2个学生自己在一个朋友圈。所以返回2。
// 
//
// 示例 2: 
//
// 
//输入: 
//[[1,1,0],
// [1,1,1],
// [0,1,1]]
//输出: 1
//说明：已知学生0和学生1互为朋友，学生1和学生2互为朋友，所以学生0和学生2也是朋友，所以他们三个在一个朋友圈，返回1。
// 
//
// 注意： 
//
// 
// N 在[1,200]的范围内。 
// 对于所有学生，有M[i][i] = 1。 
// 如果有M[i][j] = 1，则有M[j][i] = 1。 
// 
// Related Topics 深度优先搜索 并查集
package cn

//leetcode submit region begin(Prohibit modification and deletion)
func findCircleNum(M [][]int) int {
	// 1 dfs 利用条件M[i][j] = 1, 则M[j][i] = 1
	// N students
	//visited := make([]int, len(M)) // 统计第N个同学的访问状态
	//count := 0
	//for i := 0; i < len(M); i++ {
	//	if visited[i] == 0 {
	//		findDFS(M, &visited, i)
	//		count++
	//	}
	//}
	//return count

	// 2 并查集
	count := 0
	find := func(p []int, i int) int {
		root := i
		for root != p[root] {
			root = p[root]
		}
		for p[i] != i {
			x := i
			i = p[i]
			p[x] = root
		}
		return root
	}

	union := func(p []int, i, j int) {
		p1 := find(p, i)
		p2 := find(p, j)
		if p[p1] != p2 {
			p[p1] = p2
			count--
		}
	}

	// init
	d := make([]int, len(M))
	for i := 0; i < len(M); i++ {
		d[i] = i
		count++
	}

	for i := 0; i < len(M); i++ {
		for j := 0; j <= i; j++{
			if M[i][j] == 1 && i != j{
				union(d, i, j)
			}
		}
	}

	// 返回数量
	return count
}

func findDFS (M [][]int, visited *[]int, i int) {
	// 判定第i个学生和第j个学生的关系
	for j := 0; j < len(M); j++ {
		if M[i][j] == 1 && (*visited)[j] == 0 {
			(*visited)[j] = 1
			findDFS(M, visited, j)
		}
	}
}
//leetcode submit region end(Prohibit modification and deletion)
