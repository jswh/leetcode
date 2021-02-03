package leetcode
/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	current := pre.Next
	for current != nil {
		if pre.Val == current.Val {
			pre.Next = current.Next
		} else {
			pre = current
		}
		current = pre.Next
	}
	return head
}
// @lc code=end

