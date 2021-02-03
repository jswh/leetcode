/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	type Node struct {
		Level int
		TN *TreeNode
	}
	queue := []Node{}
	queue = append(queue, Node{0, root})
	for len(queue) > 0 {
		n := queue[0]
		if len(result) < n.Level + 1 {
			result = append(result, []int{})
		}
		result[n.Level] = append(result[n.Level], n.TN.Val)
		if n.TN.Left != nil {
			queue = append(queue, Node{n.Level + 1, n.TN.Left})
		}
		if n.TN.Right != nil {
			queue = append(queue, Node{n.Level + 1, n.TN.Right})
		}
		queue = queue[1:]
	}
	return result
}
// @lc code=end

