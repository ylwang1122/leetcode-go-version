//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。 
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。 
//
// 示例 1: 
//
// 输入: [3,2,3]
//输出: 3 
//
// 示例 2: 
//
// 输入: [2,2,1,1,1,2,2]
//输出: 2
// 
// Related Topics 位运算 数组 分治算法
package cn

//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	// 1. 排序
	// sort.Ints(nums)
	// return nums[len(nums)/2]

	// 2 hash
	//m := make(map[int]int)
	//n := len(nums) / 2
	//for _, v := range nums {
	//	m[v]++
	//	if m[v] > n {
	//		return v
	//	}
	//}
	//return 0

	// 3 分治
	// 整个数组的众数 简化为找寻n/2数组的众数
	return majorityele(nums, 0, len(nums) - 1)
}

func majorityele(nums []int, min, max int) int {
	if min == max {
		return nums[min]
	}

	mid := (max - min) / 2 + min
	left := majorityele(nums, min, mid)
	right := majorityele(nums, mid+1, max)

	if left == right {
		return left
	}

	// 遍历确认
	count := 0
	for i := min; i <= max; i++ {
		if nums[i] ==  left {
			count++
		}
	}
	if count > (max - min) / 2 {
		return left
	} else {
		return right
	}
}

//leetcode submit region end(Prohibit modification and deletion)
