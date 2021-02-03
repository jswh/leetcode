package leetcode

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	head, lists := getSmallOneInLists(lists)
	if head != nil {
		current := head
		for current != nil {
			current.Next, lists = getSmallOneInLists(lists)
			current = current.Next
		}
	}
	return head
}
func getSmallOneInLists(lists []*ListNode) (*ListNode, []*ListNode) {
	smallListValue := int(^uint(0) >> 1)
	smallListIndex := -1
	var smallListNode *ListNode = nil
	for i, node := range lists {
		if node != nil {
			if node.Val < smallListValue {
				smallListIndex = i
				smallListNode = node
				smallListValue = node.Val
			}

		}
	}
	if smallListIndex >= 0 {
		lists[smallListIndex] = lists[smallListIndex].Next
	}
	return smallListNode, lists
}

// @lc code=end
