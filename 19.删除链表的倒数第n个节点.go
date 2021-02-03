package leetcode

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	currentNode := head
	var nthNode *ListNode = nil
	var preNthNode *ListNode = nil
	headNthFromEnd := 1

	for currentNode != nil {
		if headNthFromEnd == n {
			nthNode = head
		} else if headNthFromEnd > n {
			preNthNode = nthNode
			nthNode = preNthNode.Next

		}
		currentNode = currentNode.Next
		headNthFromEnd++
	}
	headNthFromEnd--
	if headNthFromEnd == n {
		return head.Next
	}
	preNthNode.Next = nthNode.Next

	return head

}

// @lc code=end
