package leetcode

func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return IsSameTree(root, subRoot) || IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}
