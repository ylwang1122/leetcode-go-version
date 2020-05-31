//反转一个单链表。 
//
// 示例: 
//
// 输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL 
//
// 进阶: 
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？ 
// Related Topics 链表
package cn

type ListNode struct {
	Val int
	Next *ListNode
}
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// 1. 迭代 双指针
	//if head == nil || head.Next == nil {
	//	return  head
	//}
	//
	//var pre *ListNode
	//cur := head
	//for cur != nil {
	//	temp := cur.Next
	//	cur.Next = pre
	//	pre = cur
	//	cur = temp
	//}
	//return pre

	// 2. 递归
	if head == nil || head.Next == nil {
		return head
	}
	// 记录最后一个节点 并针对每一个节点反转操作
	p := reverseList(head.Next)
	head.Next.Next = head // 关键步骤
	head.Next = nil
	return p
}
//leetcode submit region end(Prohibit modification and deletion)
