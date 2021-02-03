package leetcode

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = nil
	head, l1, l2 = getSmallOne(l1, l2)
	if head != nil {
		current := head
		for l1 != nil || l2 != nil {
			current.Next, l1, l2 = getSmallOne(l1, l2)
			current = current.Next
		}
	}
	return head
}
func getSmallOne(l1 *ListNode, l2 *ListNode) (*ListNode, *ListNode, *ListNode) {
	if l1 == nil && l2 == nil {
		return nil, nil, nil
	}
	if l1 == nil {
		return l2, nil, l2.Next
	}
	if l2 == nil {
		return l1, l1.Next, nil
	}
	if l1.Val > l2.Val {
		return l2, l1, l2.Next
	} else {
		return l1, l1.Next, l2
	}
}

// @lc code=end
