package leetcode

func SearchRotatedSortedArray(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[l] {
			if nums[l] > target || target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if target > nums[r] || target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1
}
