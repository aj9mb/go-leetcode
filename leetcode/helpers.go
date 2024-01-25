package leetcode

func Min[T int | int16 | int32 | int64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T int | int16 | int32 | int64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func GetIndex[T comparable](arr []T, val T) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}
