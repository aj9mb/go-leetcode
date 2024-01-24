package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsValidBST(root *TreeNode) bool {
	return checkNode(root, nil, nil)
}

func checkNode(node *TreeNode, min, max *TreeNode) bool {
	if node == nil {
		return true
	}
	if min != nil && min.Val >= node.Val {
		return false
	}
	if max != nil && max.Val <= node.Val {
		return false
	}
	return checkNode(node.Left, min, node) && checkNode(node.Right, node, max)
}
