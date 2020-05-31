//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。 
//
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。 
//
// 
//
// 
//
// 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。 
//
// 
//
// 
//
// 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。 
//
// 
//
// 示例: 
//
// 输入: [2,1,5,6,2,3]
//输出: 10 
// Related Topics 栈 数组
package cn

import "container/list"

//leetcode submit region begin(Prohibit modification and deletion)
type Value struct {
	Index int
	Val int
}

func largestRectangleArea(heights []int) int {
	// 双指针 pass 案例 [5,5,1,7,1,1,5,2,7,6]
	// 栈 记录左边界，处理每一根柱子所能构成的最大面积
	//max := 0
	//
	//stack := make([]Value, 0)
	//stack = append(stack, Value{
	//	Index: -1,
	//	Val:   0,
	//})
	//
	//for i, v := range heights {
	//	// 如果元素大于栈顶元素，则入栈 否则将栈顶元素取出计算其面积
	//	if v >= stack[len(stack) - 1].Val {
	//		stack = append(stack, Value{
	//			Index: i,
	//			Val:   v,
	//		})
	//	} else {
	//		for {
	//			// 出栈计算其面积
	//			cur := stack[len(stack) - 1]
	//			stack = stack[:len(stack) - 1]
	//			area := (i - stack[len(stack) - 1].Index - 1) * cur.Val
	//			if area > max {
	//				max = area
	//			}
	//
	//			if v >= stack[len(stack) - 1].Val {
	//				stack = append(stack, Value{
	//					Index: i,
	//					Val:   v,
	//				})
	//				break
	//			}
	//		}
	//	}
	//}
	//
	//for {
	//	if len(stack) == 1 {
	//		break
	//	}
	//	cur := stack[len(stack) - 1]
	//	stack = stack[:len(stack) - 1]
	//	area := (len(heights) - stack[len(stack) - 1].Index - 1) * cur.Val
	//	if area > max {
	//		max = area
	//	}
	//}
	//return max

	// 柱状图中最大的矩形
	// 栈记录左边界，遍历找到右边界，则计算面积，并记录
	if len(heights) == 0 {
		return 0
	}

	// 双向链表模拟栈
	stack := list.New()
	stack.PushBack(Value{Index:-1, Val:-1})

	max := 0
	for i, v := range heights {
		// 依次入栈直到遇到比栈顶元素大的
		for {
			if v > stack.Back().Value.(Value).Val {
				stack.PushBack(Value{Index:i, Val:v})
				break
			} else {
				// 出栈计算面积
				cur := stack.Back()
				stack.Remove(cur)

				area := cur.Value.(Value).Val * (i - 1 - stack.Back().Value.(Value).Index)
				if area > max {
					max = area
				}
			}
		}
	}

	// 依次出栈里面的所有元素
	for ;stack.Len() != 1; {
		// 出栈计算面积
		cur := stack.Back()
		stack.Remove(cur)

		area := cur.Value.(Value).Val * (len(heights) - 1 - stack.Back().Value.(Value).Index)
		if area > max {
			max = area
		}
	}

	return max
}
//leetcode submit region end(Prohibit modification and deletion)
