//给定一个二叉树，找出其最大深度。 
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。 
//
// 说明: 叶子节点是指没有子节点的节点。 
//
// 示例： 
//给定二叉树 [3,9,20,null,null,15,7]， 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7 
//
// 返回它的最大深度 3 。 
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
func maxDepth(root *TreeNode) int {
	max := 0
	depth(root, 0, &max)
	return max
}

func depth(root *TreeNode, h int, max *int) {
	if root == nil {
		if *max < h {
			*max = h
		}
		return
	}
	depth(root.Left, h+1, max)
	depth(root.Right, h+1, max)
}
//leetcode submit region end(Prohibit modification and deletion)
