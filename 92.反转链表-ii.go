package leetcode

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	return reverseBetweenA(head, m, n)
}

func reverseBetweenA(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || n-m == 0 {
		return head
	}
	switcher := head
	switcherId := 1

	// section marks
	var preSwitcherSection *ListNode = nil
	var afterSwitcherSection *ListNode = nil
	var switcherSectionHead *ListNode = nil
	var switcherSectionTail *ListNode = nil

	// reverse temp
	var preSwitcher *ListNode = nil

	for switcher != nil {
		next := switcher.Next
		if switcherId == m-1 {
			preSwitcherSection = switcher
		} else if switcherId == m {
			switcherSectionHead = switcher
			preSwitcher = switcher
		} else if switcherId == n {
			afterSwitcherSection = switcher.Next
			switcherSectionTail = switcher
			switcher.Next = preSwitcher
		} else if switcherId > m && switcherId < n {
			switcher.Next = preSwitcher
			preSwitcher = switcher
		}

		switcher = next
		switcherId++

	}

	switcherSectionHead.Next = afterSwitcherSection
	if preSwitcherSection != nil {
		preSwitcherSection.Next = switcherSectionTail
	} else {
		head = switcherSectionTail
	}

	return head
}

// @lc code=end
