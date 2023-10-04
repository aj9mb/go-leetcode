package leetcode

func MaxSlidingWindow(nums []int, k int) []int {
	op := []int{}
	q := []int{}
	l, r := 0, 0
	for r < len(nums) {
		lq := len(q)
		for lq > 0 && nums[q[lq-1]] < nums[r] {
			q = q[:lq-1]
			lq = len(q)
		}
		q = append(q, r)
		if l > q[0] {
			q = q[1:]
		}
		if (r + 1) >= k {
			op = append(op, nums[q[0]])
			l++
		}
		r++
	}
	return op
}
