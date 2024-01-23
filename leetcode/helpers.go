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
