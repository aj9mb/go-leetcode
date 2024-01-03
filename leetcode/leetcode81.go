package leetcode

func SearchRotatedSortedArray2(nums []int, target int) bool {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		if nums[l] == nums[mid] && nums[mid] == nums[r] {
			l += 1
			r -= 1
		} else if nums[mid] >= nums[l] {
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
	return false
}
