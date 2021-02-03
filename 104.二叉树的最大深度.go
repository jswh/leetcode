package leetcode


/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(a, b int) int {
	if (a > b) {
		return a
	}
	return b
}
func maxDepth(root *TreeNode) int {
	if (root == nil) {
		return 0
	}
	if (root.Left == nil && root.Right == nil) {
		return 1
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
// @lc code=end

