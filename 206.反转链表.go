package leetcode

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	return reverseListB(head)
}
func reverseListA(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	switcher := head.Next
	preSwitcher := head
	head.Next = nil
	for switcher != nil {
		temp := switcher.Next
		switcher.Next = preSwitcher
		preSwitcher = switcher
		switcher = temp
	}
	return preSwitcher
}
func reverseListB(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	subHead := head.Next
	head.Next = nil
	newHead := reverseListB(subHead)
	subHead.Next = head

	return newHead
}

// @lc code=end
