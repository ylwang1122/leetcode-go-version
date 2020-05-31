//给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。 
//
// 示例: 
//
// 输入: [-2,1,-3,4,-1,2,1,-5,4],
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
// 
//
// 进阶: 
//
// 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。 
// Related Topics 数组 分治算法 动态规划
package cn

//leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// 动态规划 d[i] 到d[i]的最大和
	//d := make([]int, len(nums))
	//d[0] = nums[0]
	//max := d[0]
	//for i := 1; i < len(nums); i++ {
	//	d[i] = maxInt(nums[i], d[i-1]+nums[i])
	//	if d[i] > max {
	//		max = d[i]
	//	}
	//}
	//return max

	// 分治
	max := nums[0]
	maxSub(1, nums, nums[0], &max)
	return max
}

func maxSub(i int, nums []int, sum int, max *int) {
	if i == len(nums) {
		if sum > *max {
			*max = sum
		}
		return
	}

	if sum > *max {
		*max = sum
	}

	if sum + nums[i] < nums[i] {
		sum = nums[i]
	}else {
		sum += nums[i]
	}
	maxSub(i+1, nums, sum, max)
}

func maxInt (x, y int) int {
	if x > y {
		return x
	}
	return y
}
//leetcode submit region end(Prohibit modification and deletion)
