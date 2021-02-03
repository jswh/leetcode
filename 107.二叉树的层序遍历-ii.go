package leetcode
/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	type Node struct {
		Level int
		TN *TreeNode
	}
	queue := []Node{Node{0, root}}
	result := [][]int{}
	for len(queue) > 0 {
		//pop
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
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]

	}
	return result
}
// @lc code=end

