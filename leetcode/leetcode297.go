package leetcode

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func CodecConstructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (codec *Codec) Serialize(root *TreeNode) string {
	// var ser string
	// if root == nil {
	// 	return ""
	// }
	// queue := []*TreeNode{root}
	// for len(queue) > 0 {
	// 	tempQ := []*TreeNode{}
	// 	for _, node := range queue {
	// 		ser += strconv.Itoa(node.Val) + ","
	// 		if node.Left != nil {
	// 			tempQ = append(tempQ, node.Left)
	// 		} else {
	// 			ser += "null,"
	// 		}
	// 		if node.Right != nil {
	// 			tempQ = append(tempQ, node.Right)
	// 		} else {
	// 			ser += "null,"
	// 		}
	// 	}
	// 	queue = tempQ
	// }
	// if ser != "" {
	// 	ser = ser[:len(ser)-1]
	// }

	var dfs func(*TreeNode) string

	dfs = func(node *TreeNode) string {
		str := ""
		if node == nil {
			return "null"
		}
		str += strconv.Itoa(node.Val) + ","
		str += dfs(node.Left) + ","
		str += dfs(node.Right)
		return str
	}

	return dfs(root)
}

// Deserializes your encoded data to tree.
func (codec *Codec) Deserialize(data string) *TreeNode {
	splt := strings.Split(data, ",")
	i := 0
	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if splt[i] == "null" {
			i++
			return nil
		}
		val, _ := strconv.Atoi(splt[i])
		node := &TreeNode{Val: val}
		i++
		node.Left = dfs()
		node.Right = dfs()
		return node
	}
	return dfs()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
