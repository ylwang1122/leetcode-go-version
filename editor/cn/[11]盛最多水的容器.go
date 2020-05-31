//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, 
//ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。 
//
// 说明：你不能倾斜容器，且 n 的值至少为 2。 
//
// 
//
// 
//
// 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。 
//
// 
//
// 示例： 
//
// 输入：[1,8,6,2,5,4,8,3,7]
//输出：49 
// Related Topics 数组 双指针

package cn

//leetcode submit region begin(Prohibit modification and deletion)
// 暴力求解
//func maxArea(height []int) int {
//	max := 0
//	for i := 0; i < len(height); i++ {
//		for j := i; j < len(height); j++ {
//			size := (j - i) * minInt(height[i], height[j])
//			if max < size {
//				max = size
//			}
//		}
//	}
//
//	return max
//}

func maxArea(height []int) int {
	// 双指针求解 分别指向头尾，并向中间移动较小的边，直到两指针相遇
	i := 0
	max := 0
	for j := len(height) - 1; i < j; {
		size := (j - i) * minInt(height[i], height[j])
		if max < size {
			max = size
		}

		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}

	return max
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}
//leetcode submit region end(Prohibit modification and deletion)
