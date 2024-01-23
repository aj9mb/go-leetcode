package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func DiameterOfBinaryTree(root *TreeNode) int {
	max := 0
	dfs(root, &max)
	return max
}

func dfs(node *TreeNode, max *int) int {
	if node == nil {
		return -1
	}
	left := dfs(node.Left, max)
	right := dfs(node.Right, max)
	*max = Max(*max, 2+left+right)
	return 1 + Max(left, right)
}
