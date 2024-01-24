package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		mergedLists := make([]*ListNode, 0)
		for i := 0; i < len(lists); i += 2 {
			l1 := lists[i]
			var l2 *ListNode
			if (i + 1) < len(lists) {
				l2 = lists[i+1]
			}
			mergedLists = append(mergedLists, mergeLists(l1, l2))
		}
		lists = mergedLists
	}
	return lists[0]
}

func mergeLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := ListNode{}
	tail := &dummy
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			tail.Next = l2
			l2 = l2.Next
		} else {
			tail.Next = l1
			l1 = l1.Next
		}
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	return dummy.Next
}
