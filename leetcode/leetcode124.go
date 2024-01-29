package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func MaxPathSum(root *TreeNode) int {
	max := int(math.Inf(-1))

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMax := Max(dfs(node.Left), 0)
		rightMax := Max(dfs(node.Right), 0)
		max = Max(max, rightMax+leftMax+node.Val)
		return Max(leftMax, rightMax) + node.Val
	}

	dfs(root)
	return max
}
