package main

import (
	"fmt"

	"github.com/aj9mb/go-leetcode/leetcode"
)

func main() {
	// fmt.Print(leetcode.MinWindow("ADOBECODEBANC", "ABC"))
	// fmt.Print(leetcode.MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	// fmt.Print(leetcode.Search([]int{-1, 0, 3, 5, 9, 12}, 2))
	// fmt.Print(leetcode.MinEatingSpeed([]int{312884470}, 968709470))
	// fmt.Print(leetcode.FindMin([]int{3, 1, 2}))
	// fmt.Print(leetcode.SearchRotatedSortedArray2([]int{1, 2, 1}, 1))
	// fmt.Print(leetcode.FindMedianSortedArrays([]int{}, []int{1}))
	// fmt.Print(leetcode.GenerateParanthesis(1))

	// arrToLL := func(op []int) *leetcode.ListNode {
	// 	head := new(leetcode.ListNode)

	// 	temp := head

	// 	for i, v := range op {
	// 		temp.Val = v
	// 		if i < len(op)-1 {
	// 			temp.Next = new(leetcode.ListNode)
	// 			temp = temp.Next
	// 		}
	// 	}
	// 	return head
	// }

	// op1 := []int{1, 2, 3, 4, 5}
	// head1 := arrToLL(op1)

	// op2 := []int{5, 6, 4}
	// head2 := arrToLL(op2)

	// res := leetcode.AddTwoNumbers(head1, head2)
	// temp := res
	// for temp != nil {
	// 	fmt.Printf("%d, ", temp.Val)
	// 	temp = temp.Next
	// }

	// fmt.Print(leetcode.FindDuplicate([]int{1, 3, 4, 2, 2}))

	// res := leetcode.ReverseKGroup(head1, 2)
	// temp := res
	// for temp != nil {
	// 	fmt.Printf("%d, ", temp.Val)
	// 	temp = temp.Next
	// }

	// root := &leetcode.TreeNode{Val: 3, Left: &leetcode.TreeNode{Val: 9}, Right: &leetcode.TreeNode{Val: 20, Left: &leetcode.TreeNode{Val: 15}, Right: &leetcode.TreeNode{Val: 7}}}

	// root := &leetcode.TreeNode{Val: -10, Left: &leetcode.TreeNode{Val: 9}, Right: &leetcode.TreeNode{Val: 20, Left: &leetcode.TreeNode{Val: 15}, Right: &leetcode.TreeNode{Val: 7}}}
	// root := &leetcode.TreeNode{Val: -3}
	// fmt.Print(leetcode.MaxPathSum(root))
	// codec := leetcode.CodecConstructor()
	// fmt.Print(codec.Serialize(root))

	fmt.Print(leetcode.Subsets([]int{1, 2, 3}))
}
