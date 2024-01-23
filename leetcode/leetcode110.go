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
func IsBalanced(root *TreeNode) bool {
	balanced, _ := getHeightAndCheckIsBalanced(root)
	return balanced
}

func getHeightAndCheckIsBalanced(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}
	leftBalanced, leftHeight := getHeightAndCheckIsBalanced(node.Left)
	rightBalanced, rightHeight := getHeightAndCheckIsBalanced(node.Right)

	balanced := rightBalanced && leftBalanced && (math.Abs(float64(leftHeight)-float64(rightHeight)) <= 1)
	return balanced, (1 + Max(rightHeight, leftHeight))
}
