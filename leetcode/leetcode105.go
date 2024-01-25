package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 || len(inorder) < 1 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	mid := GetIndex(inorder, preorder[0])

	root.Left = BuildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = BuildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}
