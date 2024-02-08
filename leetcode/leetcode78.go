package leetcode

func Subsets(nums []int) [][]int {
	result := make([][]int, 0)
	subset := make([]int, 0)
	var backtrack func(i int)
	backtrack = func(i int) {
		if i >= len(nums) {
			result = append(result, append([]int{}, subset...))
			return
		}

		subset = append(subset, nums[i])
		backtrack(i + 1)

		subset = subset[:len(subset)-1]
		backtrack(i + 1)
	}

	backtrack(0)
	return result
}
