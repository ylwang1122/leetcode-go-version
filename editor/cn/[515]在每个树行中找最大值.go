//您需要在二叉树的每一行中找到最大的值。 
//
// 示例： 
//
// 
//输入: 
//
//          1
//         / \
//        3   2
//       / \   \  
//      5   3   9 
//
//输出: [1, 3, 9]
// 
// Related Topics 树 深度优先搜索 广度优先搜索

package cn

import (
	"container/list"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	// 1.bfs
	if root == nil {
		return nil
	}

	queue1 := list.New()
	queue2 := list.New()
	queue1.PushBack(root)

	ret := make([]int, 0)
	for queue1.Len()  != 0 {
		flag := true
		min := 0
		for queue1.Len() != 0 {
			front := queue1.Front()
			queue1.Remove(front)
			node :=  front.Value.(*TreeNode)
			if flag || min < node.Val {
				min = node.Val
				flag = false
			}
			if node.Left != nil {
				queue2.PushBack(node.Left)
			}
			if node.Right != nil {
				queue2.PushBack(node.Right)
			}
		}
		ret = append(ret, min)
		queue1, queue2 = queue2, queue1
	}
	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
