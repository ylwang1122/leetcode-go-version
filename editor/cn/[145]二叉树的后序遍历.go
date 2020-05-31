//给定一个二叉树，返回它的 后序 遍历。 
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
//输出: [3,2,1] 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树
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
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := list.New()
	ret := make([]int, 0)
	stack.PushBack(root)
	cur := root
	for ; stack.Len() != 0; {
		top := stack.Back()
		if top.Value.(*TreeNode).Left != nil && top.Value.(*TreeNode).Left != cur && top.Value.(*TreeNode).Right != cur {
			stack.PushBack(top.Value.(*TreeNode).Left)
		} else if top.Value.(*TreeNode).Right != nil && top.Value.(*TreeNode).Right != cur {
			stack.PushBack(top.Value.(*TreeNode).Right)
		} else {
			stack.Remove(top)
			ret = append(ret, top.Value.(*TreeNode).Val)
			cur = top.Value.(*TreeNode)
		}
	}

	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
