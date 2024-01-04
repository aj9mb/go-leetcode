package leetcode

func copyRandomList(head *Node) *Node {
	lookup := make(map[*Node]*Node)
	temp := head
	for temp != nil {
		newNode := new(Node)
		newNode.Val = temp.Val
		lookup[temp] = newNode
		temp = temp.Next
	}
	temp = head
	for temp != nil {
		newNode := lookup[temp]
		newNode.Next = lookup[temp.Next]
		newNode.Random = lookup[temp.Random]
		temp = temp.Next
	}
	return lookup[head]
}
