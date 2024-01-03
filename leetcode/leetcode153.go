package leetcode

func FindMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	min := nums[0]
	for l <= r {
		if nums[l] <= nums[r] {
			if min > nums[l] {
				min = nums[l]
			}
			return min
		}
		mid := (l + r) / 2
		if min > nums[mid] {
			min = nums[mid]
		}
		if nums[mid] >= nums[l] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return min
}
