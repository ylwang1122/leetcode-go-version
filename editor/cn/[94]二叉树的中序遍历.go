//给定一个二叉树，返回它的中序 遍历。 
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
//输出: [1,3,2] 
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？ 
// Related Topics 栈 树 哈希表
package cn

import "container/list"

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
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	// 1.遍历
	//ret := make([]int, 0, 3)
	//ret = append(ret, inorderTraversal(root.Left)...)
	//// 输出
	//ret = append(ret, root.Val)
	//ret = append(ret, inorderTraversal(root.Right)...)
	//return ret

	// 2.栈
	stack := list.New()

	// 左子树 输出 右子树
	ret := make([]int, 0)
	cur := root
	for ; cur != nil || stack.Len() != 0; {
		for ; cur != nil; {
			stack.PushBack(cur)
			cur = cur.Left
		}
		// 出栈
		b := stack.Back()
		stack.Remove(b)
		ret = append(ret, b.Value.(*TreeNode).Val)
		cur = b.Value.(*TreeNode).Right
	}
	return ret
}
//leetcode submit region end(Prohibit modification and deletion)
