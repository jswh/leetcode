package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	flag := 0
	var head *ListNode
	var current *ListNode
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += flag
		flag = 0
		if sum > 9 {
			flag = 1
			sum -= 10
		}
		node := ListNode{sum, nil}
		if head == nil {
			head = &node
			current = head
		} else {
			current.Next = &node
			current = current.Next
		}
	}
	if flag == 1 {
		node := ListNode{1, nil}
		current.Next = &node
	}
	return head
}

// @lc code=end
