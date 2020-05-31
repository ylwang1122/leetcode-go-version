//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
// 使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。 
//
// 注意：答案中不可以包含重复的三元组。 
//
// 
//
// 示例： 
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
// 
// Related Topics 数组 双指针
package cn

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	// 双指针，转化为两数之和的问题
	sort.Ints(nums)

	// 如何去重
	ret := make([][]int, 0)
	for k := 0; k < len(nums) - 2; k++ {
		// 优化 j > i > k不可能组成之和为0
		if nums[k] > 0 {
			break
		}
		// 去重
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		target := -nums[k]
		i := k + 1
		for j := len(nums) - 1; j > i; {
			sum := nums[i] + nums[j]
			if sum == target {
				ret = append(ret, []int{nums[k], nums[i], nums[j]})
				// 去重
				for i++;j > i; i++ {
					if nums[i] != nums[i-1] {
						break
					}
				}
				for j-- ;j > i; j-- {
					if nums[j] != nums[j+1] {
						break
					}
				}
			} else if sum > target {
				for j-- ;j > i; j-- {
					if nums[j] != nums[j+1] {
						break
					}
				}
			} else {
				for i++;j > i; i++ {
					if nums[i] != nums[i-1] {
						break
					}
				}
			}
		}
	}

	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
