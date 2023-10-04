package leetcode

func SearchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		if row[len(row)-1] < target {
			continue
		}
		start, end := 0, len(row)-1
		for start <= end {
			mid := (start + end) / 2
			if row[mid] == target {
				return true
			} else if row[mid] < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
		if row[len(row)-1] > target {
			return false
		}
	}
	return false
}
