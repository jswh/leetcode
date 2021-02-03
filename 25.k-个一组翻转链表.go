package leetcode

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	var kthNode *ListNode = head
	i := 1
	for i < k && kthNode.Next != nil {
		i++
		kthNode = kthNode.Next
	}
	if i < k {
		return head
	}
	nextGroupHead := kthNode.Next
	nextGroupHead = reverseKGroup(nextGroupHead, k)
	kthNode.Next = nil
	newHead := doReverseList(head)
	head.Next = nextGroupHead

	return newHead
}

func doReverseList(head *ListNode) *ListNode {
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

// @lc code=end
