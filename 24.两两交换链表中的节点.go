package leetcode

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	//至少有两个才能交换
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	prePairNode := head
	doSwap(head, head.Next)

	var pairA *ListNode = nil
	var pairB *ListNode = nil
	for prePairNode.Next != nil && prePairNode.Next.Next != nil {
		pairA = prePairNode.Next
		pairB = pairA.Next
		doSwap(pairA, pairB)
		prePairNode.Next = pairB

		prePairNode = pairA
	}

	return newHead
}
func doSwap(pairA *ListNode, pairB *ListNode) {
	afterPairNode := pairB.Next
	pairB.Next = pairA
	pairA.Next = afterPairNode
}

// @lc code=end
