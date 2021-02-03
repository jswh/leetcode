package leetcode


/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	// 基于树的镜像
	if root == nil {
		return true
	}
	// 另可以bfs的方法，按层取值之后看是否对称
	return isMirror(root.Left, root.Right)
}
func isMirror(a *TreeNode, b *TreeNode) bool {
	// 利用的是互为镜像的树的性质，
	// a, b 树互为镜像，表达为a <-> b则
	// 1.节点值相同
	// 2 a.Left <-> b.Rihgt, a.Right <-> b.Left
	if a == nil && b == nil{
		return true
	}
	if (a == nil && b != nil) || (a != nil && b == nil) {
		return false
	}
	return a.Val == b.Val && isMirror(a.Left, b.Right) && isMirror(a.Right, b.Left)
}
// @lc code=end

