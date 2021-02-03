package leetcode

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func isValidBST(root *TreeNode) bool {
	list := inorderTraversal(root)
	pre := MinInt
	for _, v := range list {
		if pre >= v {
			return false
		} else {
			pre = v
		}
	}
	return true
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return make([]int, 0)
	}
	var left []int
	var right []int
	if root.Left != nil {
		left = inorderTraversal(root.Left)
	}
	if root.Right != nil {
		right = inorderTraversal(root.Right)
	}
	result := []int{}
	if left != nil {
		result = append(result, left...)
	}
	result = append(result, root.Val)
	if right != nil {
		result = append(result, right...)
	}

	return result
}

// @lc code=end
