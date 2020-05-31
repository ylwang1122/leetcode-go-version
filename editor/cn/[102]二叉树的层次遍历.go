//给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。 
//
// 例如: 
//给定二叉树: [3,9,20,null,null,15,7], 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
// 
//
// 返回其层次遍历结果： 
//
// [
//  [3],
//  [9,20],
//  [15,7]
//]
// 
// Related Topics 树 广度优先搜索
package cn

import "container/list"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type nodeInfo struct {
	node *TreeNode
	height int
}

func levelOrder(root *TreeNode) [][]int {
	// 3 bfs 广度优先遍历 双队列
	if root == nil {
		return nil
	}

	queue1 := list.New()
	queue2 := list.New()
	queue1.PushBack(root)

	rets := make([][]int, 0)
	for queue1.Len() != 0 {
		ret := make([]int, 0)
		for queue1.Len() != 0 {
			front := queue1.Front()
			queue1.Remove(front)
			ret = append(ret, front.Value.(*TreeNode).Val)
			if front.Value.(*TreeNode).Left != nil {
				queue2.PushBack(front.Value.(*TreeNode).Left)
			}
			if front.Value.(*TreeNode).Right != nil {
				queue2.PushBack(front.Value.(*TreeNode).Right)
			}
		}
		rets = append(rets, ret)
		queue1, queue2 = queue2, queue1
	}

	return rets
}

func dfsH(height int, root *TreeNode, output map[int][]int) {
	if root == nil {
		return
	}

	// 记录数据
	output[height] = append(output[height], root.Val)

	dfsH(height + 1, root.Left, output)
	dfsH(height + 1, root.Right, output)
}
//leetcode submit region end(Prohibit modification and deletion)
