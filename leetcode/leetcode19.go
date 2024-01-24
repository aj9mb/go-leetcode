package leetcode

func RmoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head, Val: 0}
	left := dummy
	right := head
	for i := 0; i < n; i++ {
		right = right.Next
	}
	for right != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return dummy.Next
}
