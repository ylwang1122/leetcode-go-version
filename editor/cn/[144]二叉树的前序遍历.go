//给定一个二叉树，返回它的 前序 遍历。 
//
// 示例: 
//
// 输入: [1,null,2,3]  
//   1
//    \
//     2
//    /
//   3 
//
//输出: [1,2,3]
// 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树

package cn

import (
	"container/list"
)

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := list.New()
	ret := make([]int, 0)
	cur := root
	for ;cur != nil || stack.Len() != 0; {
		for ; cur != nil; {
			ret = append(ret, cur.Val)
			stack.PushBack(cur)
			cur = cur.Left
		}
		b := stack.Back()
		stack.Remove(b)
		cur = b.Value.(*TreeNode).Right
	}
	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
