package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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

func inorderTraversal(root *TreeNode) []int {
	//return inorderTraversal2(root)
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

type Stack []*TreeNode

//Push push an element to the stack
func (s *Stack) Push(item *TreeNode) {
	*s = append(*s, item)
}

//IsEmpty check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

//Pop pop an element from the stack
func (s *Stack) Pop() *TreeNode {
	if s.IsEmpty() {
		return nil
	}
	index := len(*s) - 1
	item := (*s)[index]
	*s = (*s)[:index]
	return item
}

func inorderTraversal2(root *TreeNode) []int {
	result := []int{}
	stack := Stack{}
	current := root
	//栈循环的本质其实时手动模拟了递归过程
	for current != nil || !stack.IsEmpty() {
		//持续访问左子树，直到左子树为空
		if current != nil {
			stack.Push(current)
			current = current.Left
		} else {
			//左子树为空时回退到父节点，添加结果，开始访问右子树
			//右子树也为空的话，继续回退到父节点
			current = stack.Pop()
			result = append(result, current.Val)
			current = current.Right
		}
	}

	return result
}

// @lc code=end
func main() {
	l := TreeNode{Val: 3, Left: nil, Right: nil}
	right := TreeNode{
		Val:   2,
		Left:  &l,
		Right: nil,
	}
	root := TreeNode{
		Val:   1,
		Left:  nil,
		Right: &right,
	}
	inorderTraversal2(&root)
}
