package leetcode

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k < 1 {
		return head
	}
	tail := head
	lenght := 1
	for tail.Next != nil {
		tail = tail.Next
		lenght++
	}
	//先找到链表尾部，右旋k，本质是将倒数K作为新的head，然后把原来的tail.Next指向head
	k = k % lenght
	tail.Next = head
	for i := 0; i < (lenght - k - 1); i++ {
		head = head.Next
	}
	newHead := head.Next
	head.Next = nil

	return newHead
}

// @lc code=end
