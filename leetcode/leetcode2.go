package leetcode

import "strconv"

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Temp := l1
	l2Temp := l2
	res := &ListNode{}
	resCurr := res
	carry := 0
	for l1Temp != nil || l2Temp != nil || carry > 0 {
		k := 0
		if l1Temp != nil {
			k += l1Temp.Val
			l1Temp = l1Temp.Next
		}
		if l2Temp != nil {
			k += l2Temp.Val
			l2Temp = l2Temp.Next
		}
		k += carry
		if k > 9 {
			carry = k / 10
			k = k % 10
		} else {
			carry = 0
		}
		resCurr.Val = k
		if l1Temp != nil || l2Temp != nil || carry > 0 {
			resCurr.Next = &ListNode{}
			resCurr = resCurr.Next
		}
	}
	return res
}

func Add2Numbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Arr := make([]int, 0)
	l1Curr := l1
	for l1Curr != nil {
		l1Arr = append(l1Arr, l1Curr.Val)
		l1Curr = l1Curr.Next
	}

	l2Arr := make([]int, 0)
	l2Curr := l2
	for l2Curr != nil {
		l2Arr = append(l2Arr, l2Curr.Val)
		l2Curr = l2Curr.Next
	}

	l1Str := ""
	for i := len(l1Arr) - 1; i >= 0; i-- {
		l1Str += strconv.Itoa(l1Arr[i])
	}

	var l1Int int64
	l1Int, _ = strconv.ParseInt(l1Str, 10, 64)

	l2Str := ""
	for i := len(l2Arr) - 1; i >= 0; i-- {
		l2Str += strconv.Itoa(l2Arr[i])
	}
	var l2Int int64
	l2Int, _ = strconv.ParseInt(l2Str, 10, 64)

	opInt := l1Int + l2Int

	opStr := strconv.FormatInt(opInt, 10)

	op := make([]int, 0)
	for i := len(opStr) - 1; i >= 0; i-- {
		in, _ := strconv.Atoi(string(opStr[i]))
		op = append(op, in)
	}

	head := new(ListNode)

	temp := head

	for i, v := range op {
		temp.Val = v
		if i < len(op)-1 {
			temp.Next = new(ListNode)
			temp = temp.Next
		}
	}

	return head
}
