//假设按照升序排序的数组在预先未知的某个点上进行了旋转。 
//
// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。 
//
// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。 
//
// 你可以假设数组中不存在重复的元素。 
//
// 你的算法时间复杂度必须是 O(log n) 级别。 
//
// 示例 1: 
//
// 输入: nums = [4,5,6,7,0,1,2], target = 0
//输出: 4
// 
//
// 示例 2: 
//
// 输入: nums = [4,5,6,7,0,1,2], target = 3
//输出: -1 
// Related Topics 数组 二分查找
package cn

//leetcode submit region begin(Prohibit modification and deletion)
func search(nums []int, target int) int {
	// 二分查找
	// 审题 元素中没有重复元素 判定那侧是连续
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { // 左侧连续
			if nums[mid] >= target && target >= nums[left]{ // 判定是否位于左侧连续序列中
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右侧连续
			if nums[mid] <= target && target <= nums[right]{ // 判定是否位于右侧连续序列中
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
//leetcode submit region end(Prohibit modification and deletion)
