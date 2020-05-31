//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。 
//
// 示例 1: 
//
// 输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
// 
//
// 示例 2: 
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4 
//
// 说明: 
//
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。 
// Related Topics 堆 分治算法
package  cn

//leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	for i := k / 2 - 1; i >= 0; i-- {
		adjustHeap(nums, i, k)
	}

	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0] = nums[i]
			adjustHeap(nums, 0, k)
		}
	}
	return nums[0]
}

// 构建最小堆 堆容量为k 堆顶即为所求
func adjustHeap(nums []int, i, len int) {
	temp := nums[i]
	for k := 2 * i + 1; k < len; k = 2*k + 1 {
		if k + 1 < len && nums[k] > nums[k+1] {
			k++
		}

		if nums[k] < temp {
			nums[i] = nums[k]
			i = k
		} else {
			break
		}
		nums[i] = temp
	}
}
//leetcode submit region end(Prohibit modification and deletion)
