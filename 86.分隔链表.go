package leetcode

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func partition(head *ListNode, x int) *ListNode {
	var smallHead *ListNode = nil
	var smallPoint *ListNode = nil
	var largeHead *ListNode = nil
	var largePoint *ListNode = nil
	current := head
	for current != nil {
		if current.Val < x {
			if smallHead == nil {
				smallHead = current
				smallPoint = current
			} else {
				smallPoint.Next = current
				smallPoint = current
			}
		} else {
			if largeHead == nil {
				largeHead = current
				largePoint = current
			} else {
				largePoint.Next = current
				largePoint = current
			}

		}
		current = current.Next
	}
	if smallPoint != nil {
		smallPoint.Next = largeHead
	}
	if largePoint != nil {
		largePoint.Next = nil
	}
	if smallHead != nil {
		return smallHead
	} else {
		return largeHead
	}

}

// @lc code=end
