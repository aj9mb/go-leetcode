package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func GoodNodes(root *TreeNode) int {
	count := 0
	if root != nil {
		count = checkGoodNodes(root, root.Val)
	}
	return count
}

func checkGoodNodes(curr *TreeNode, max int) int {
	if curr == nil {
		return 0
	}
	res := 0
	if max <= curr.Val {
		res++
		max = curr.Val
	}
	res += checkGoodNodes(curr.Left, max)
	res += checkGoodNodes(curr.Right, max)
	return res
}
