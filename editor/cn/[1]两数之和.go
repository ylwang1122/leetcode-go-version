//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。 
//
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。 
//
// 示例: 
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
// 
// Related Topics 数组 哈希表

package cn

import "sort"

//leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	// hash表
	numMap := make(map[int]int)
	for index, val := range nums {
		numMap[val] = index
	}

	for index, value := range nums {
		if val, ok := numMap[target - value]; ok && index != val{
			return []int{index, val}
		}
	}
	return nil

	// 双指针 不可以 排序会改变原有下标的顺序
	//sort.Ints(nums)
	//i := 0
	//for j := len(nums) - 1; j > i; {
	//	sum := nums[i] + nums[j]
	//	if sum == target {
	//		return []int{i, j}
	//	} else if sum > target {
	//		j--
	//	} else {
	//		i++
	//	}
	//}
	return nil
}
//leetcode submit region end(Prohibit modification and deletion)
