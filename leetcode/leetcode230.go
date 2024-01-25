package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func KthSmallestIter(root *TreeNode, k int) int {
	n := 0
	stack := []*TreeNode{}
	curr := root

	for curr != nil || stack != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		n++
		if len(stack) == 1 {
			stack = []*TreeNode{}
		} else {
			stack = stack[:len(stack)-1]
		}
		if n == k {
			return curr.Val
		}
		curr = curr.Right
	}
	return -1
}
