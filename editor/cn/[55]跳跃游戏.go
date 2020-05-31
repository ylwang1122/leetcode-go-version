//给定一个非负整数数组，你最初位于数组的第一个位置。 
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。 
//
// 判断你是否能够到达最后一个位置。 
//
// 示例 1: 
//
// 输入: [2,3,1,1,4]
//输出: true
//解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
// 
//
// 示例 2: 
//
// 输入: [3,2,1,0,4]
//输出: false
//解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
// 
// Related Topics 贪心算法 数组

package cn
//leetcode submit region begin(Prohibit modification and deletion)
func canJump(nums []int) bool {
	// 贪心算法 每次跳跃都选择本次加下一次能跳的最远距离进行跳跃 可求得最小跳跃步数
	//for i := 0; i < len(nums); {
	//	max := 0
	//	index := i //记录本次跳跃所选节点的坐标
	//	for j := 1; j <= nums[i] && i + j < len(nums); j++ {
	//		if j + nums[i+j] > max {
	//			max = j + nums[i+j]
	//			index = i + j
	//		}
	//	}
	//	// 判定是否可以到达最后
	//	if index + nums[index] >= len(nums) || i == len(nums) - 1 {
	//		return true
	//	}
	//	if i == index {
	//		return false
	//	}
	//	i = index
	//}
	//return false

	if len(nums) < 2 {
		return true
	}

	index := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] + i >= index {
			index = i
		}
	}
	return index == 0
}
//leetcode submit region end(Prohibit modification and deletion)
