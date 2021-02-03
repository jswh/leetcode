package leetcode

/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	for head != nil && head.Next != nil && head.Val == head.Next.Val {
		head = seekNewSection(head, head.Val)
	}
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	current := pre.Next
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current = seekNewSection(current, current.Val)
			pre.Next = current
		} else {
			pre = current
			current = pre.Next
		}

	}

	return head
}
func seekNewSection(head *ListNode, delval int) *ListNode {
	for head != nil && head.Val == delval {
		head = head.Next
	}
	return head
}

// @lc code=end
