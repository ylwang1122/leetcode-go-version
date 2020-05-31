//给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。 
//
// 说明：解集不能包含重复的子集。 
//
// 示例: 
//
// 输入: nums = [1,2,3]
//输出:
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//] 
// Related Topics 位运算 数组 回溯算法

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	// 1. 递归 类比有效括号的题目，针对某一项 针对某一项数字 选或者不选
	if len(nums) == 0 {
		return [][]int{{}}
	}

	ret := make([][]int, 0)
	cur := make([]int, 0)
	dfs(nums, 0, cur, &ret)
	return ret

	// 2.数学方式
	//ret := [][]int{{}}
	//for _, num := range nums {
	//	sub := make([][]int, 0)
	//	for _, v := range ret {
	//		d := make([]int, len(v))
	//		copy(d, v)
	//		newSub := append(d, num)
	//		sub = append(sub, newSub)
	//	}
	//	ret = append(ret, sub...)
	//}
	//return ret
}

func dfs(nums []int, index int, cur []int, ret *[][]int) {
	if index == len(nums) {
		// 注意slice的传递操作，以及赋值操作，如果将其放到某一个空间中，再取出其副本时，需要使用深拷贝
		// 保证数据的正确性
		d := make([]int, len(cur))
		copy(d, cur)
		*ret = append(*ret, d)
		return
	}

	// 不选当前值
	dfs(nums, index+1, cur, ret)

	// 选当前值
	cur = append(cur, nums[index])
	dfs(nums, index+1, cur, ret)
}

//leetcode submit region end(Prohibit modification and deletion)
