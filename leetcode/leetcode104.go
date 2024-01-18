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
func MaxDepthRecurDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(MaxDepthRecurDFS(root.Left)), float64(MaxDepthRecurDFS(root.Right))))
}

func MaxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	level := 0
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		level++
	}
	return level
}

// func MaxDepthIterDFS(root *TreeNode) int { NOT working
// 	if root == nil {
// 		return 0
// 	}
// 	level := 0
// 	stack := make([]*TreeNode, 0)
// 	stack = append(stack, root)
// 	for len(stack) > 0 {
// 		l := len(stack)
// 		for i := 0; i < l; i++ {
// 			node := stack[0]
// 			stack = stack[1:]
// 			if node.Right != nil {
// 				tmp := make([]*TreeNode, 0)
// 				tmp = append(tmp, node.Right)
// 				tmp = append(tmp, stack...)
// 				stack = tmp
// 			}
// 			if node.Left != nil {
// 				tmp := make([]*TreeNode, 0)
// 				tmp = append(tmp, node.Left)
// 				tmp = append(tmp, stack...)
// 				stack = tmp
// 			}
// 		}
// 		level++
// 	}
// 	return level
// }
