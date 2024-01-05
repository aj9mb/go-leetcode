package leetcode

func FindDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]

	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}

	slow2 := 0

	for {
		slow, slow2 = nums[slow], nums[slow2]
		if slow == slow2 {
			return slow
		}
	}
}
