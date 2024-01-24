package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else {
			return root
		}
	}
	return root
}

func LowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if p.Val > root.Val && q.Val > root.Val {
		return LowestCommonAncestor1(root.Right, p, q)
	}
	if p.Val < root.Val && q.Val < root.Val {
		return LowestCommonAncestor1(root.Left, p, q)
	}
	return root
}
