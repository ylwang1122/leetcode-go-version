//给定一个二叉树，判断其是否是一个有效的二叉搜索树。 
//
// 假设一个二叉搜索树具有如下特征： 
//
// 
// 节点的左子树只包含小于当前节点的数。 
// 节点的右子树只包含大于当前节点的数。 
// 所有左子树和右子树自身必须也是二叉搜索树。 
// 二叉搜索树的中序遍历是升序序列
//
// 示例 1: 
//
// 输入:
//    2
//   / \
//  1   3
//输出: true
// 
//
// 示例 2: 
//
// 输入:
//    5
//   / \
//  1   4
//     / \
//    3   6
//输出: false
//解释: 输入为: [5,1,4,null,null,3,6]。
//     根节点的值为 5 ，但是其右子节点值为 4 。
// 
// Related Topics 树 深度优先搜索
package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	// 1. 中序遍历升序的特性
	//if root == nil {
	//	return true
	//}
	//
	//stack := list.New()
	//cur := root
	//num := 0
	//flag := false
	//for cur != nil || stack.Len() != 0 {
	//	for cur != nil {
	//		stack.PushBack(cur)
	//		cur = cur.Left
	//	}
	//
	//	top := stack.Back()
	//	stack.Remove(top)
	//	if flag && num >= top.Value.(*TreeNode).Val {
	//		return false
	//	}
	//	num = top.Value.(*TreeNode).Val
	//	flag = true
	//
	//	cur = top.Value.(*TreeNode).Right
	//}
	//return true

	// 2 递归
	return isBST(root, nil, nil)
}

func isBST(cur, min, max *TreeNode) bool {
	if cur == nil {
		return true
	}

	// 判定是否满足条件
	if min != nil && cur.Val <= min.Val {
		return false
	}
	if max != nil && cur.Val >= max.Val {
		return false
	}
	return isBST(cur.Left, min, cur) && isBST(cur.Right, cur, max)
}
//leetcode submit region end(Prohibit modification and deletion)
