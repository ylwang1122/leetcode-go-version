//给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。 
//
// 例如，给定三角形： 
//
// [
//     [2],
//    [3,4],
//   [6,5,7],
//  [4,1,8,3]
//]
// 
//
// 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。 
//
// 说明： 
//
// 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。 
// Related Topics 数组 动态规划

package cn
//leetcode submit region begin(Prohibit modification and deletion)
func minimumTotal(triangle [][]int) int {
	// 动态规划
	// d[i][j] = min(d[i-1][j], d[i-1][j-1}) + triangle[i][j]
	if len(triangle) == 0 {
		return 0
	}

	//d := make([][]int, len(triangle))
	//for i := 0; i < len(d); i++ {
	//	d[i] = make([]int, len(triangle[i]))
	//}
	//d[0][0] = triangle[0][0]
	//for i := 1; i < len(triangle); i++ {
	//	for j := 0; j < len(triangle[i]); j++ {
	//		if i > j && j != 0{
	//			d[i][j] = minInt(d[i-1][j], d[i-1][j-1]) + triangle[i][j]
	//		} else if j == 0 {
	//			d[i][j] = d[i-1][j] + triangle[i][j]
	//		} else {
	//			d[i][j] = d[i-1][j-1] + triangle[i][j]
	//		}
	//	}
	//}
	//
	//min := d[len(d) -1][0]
	//for _, v := range d[len(d) - 1] {
	//	if min > v {
	//		min = v
	//	}
	//}

	d := make([]int, len(triangle))
	d[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		var tmp int;
		pre := d[0]
		for j := 0; j < len(triangle[i]); j++ {
			tmp = d[j]
			if i > j && j != 0{
				d[j] = minInt(d[j], pre) + triangle[i][j]
			} else if j == 0 {
				d[j] = d[j] + triangle[i][j]
			} else {
				d[j] = pre + triangle[i][j]
			}
			pre = tmp
		}
	}

	min := d[0]
	for _, v := range d {
		if min > v {
			min = v
		}
	}

	return min
}

func minInt (x, y int) int {
	if x > y {
		return y
	}
	return x
}
//leetcode submit region end(Prohibit modification and deletion)
