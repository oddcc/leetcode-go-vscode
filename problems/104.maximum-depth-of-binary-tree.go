package problems
/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] Maximum Depth of Binary Tree
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	lD := maxDepth(root.Left)
	rD := maxDepth(root.Right)
	if lD >= rD {
		return lD + 1
	}
	return rD + 1
}
// @lc code=end

