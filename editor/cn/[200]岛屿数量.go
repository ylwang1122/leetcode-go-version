//给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设
//网格的四个边均被水包围。 
//
// 示例 1: 
//
// 输入:
//11110
//11010
//11000
//00000
//
//输出: 1
// 
//
// 示例 2: 
//
// 输入:
//11000
//11000
//00100
//00011
//
//输出: 3
// 
// Related Topics 深度优先搜索 广度优先搜索 并查集
package cn


//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	// 1 深度优先搜索
	//visited := make([][]int, len(grid))
	//for i := 0; i < len(grid); i++ {
	//	visited[i] = make([]int, len(grid[0]))
	//}
	//
	//count := 0
	//for i := 0; i < len(grid); i++ {
	//	for j := 0; j < len(grid[0]); j++ {
	//		if visited[i][j] == 0 && grid[i][j] == '1' {
	//			numDFS(grid, &visited, i, j)
	//			count++
	//		}
	//	}
	//}
	//return count

	// 2 并查集
	// 初始化 合并 查找顶级节点（路径压缩）
	count := 0

	find := func(p []int, i int) int {
		root := i
		for root != p[root] {
			root = p[root]
		}
		// 路径压缩
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

	// 遍历岛屿构造并查集
	// init
	if len(grid) == 0 {
		return 0
	}

	d := make([]int, len(grid) * len(grid[0])) //  每个岛看做一个单独的个体 0不看做个体
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				d[i * len(grid[0]) + j] = i * len(grid[0]) + j
				count++
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				// 判定需要合并那些岛屿
				grid[i][j] = '0'
				cur := i * len(grid[0]) + j
				// 判断右
				if j + 1 < len(grid[0]) && grid[i][j+1] == '1' {
					union(d, cur, i * len(grid[0]) + j + 1)
				}
				// 判断下
				if i + 1 < len(grid) && grid[i+1][j] == '1' {
					union(d, cur, (i+1)*len(grid[0]) + j)
				}
			}
		}
	}

	// 遍历并查集 统计岛屿数量
	return count
}

// i 表示行数
func numDFS(grid [][]byte, visited *([][]int), i, j int) {
	if  i == len(grid) || j == len(grid[0]) || i < 0 || j < 0{
		return
	}

	if grid[i][j] == '1' && (*visited)[i][j] == 0 {
		(*visited)[i][j] = 1
		numDFS(grid, visited, i+1, j)
		numDFS(grid, visited, i, j+1)
		numDFS(grid, visited, i-1, j)
		numDFS(grid, visited, i, j-1)
	}
}
//leetcode submit region end(Prohibit modification and deletion)
