package leetcode

import "math"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2
	total := len(A) + len(B)
	half := (total + 1) / 2
	if len(A) > len(B) {
		A, B = B, A
	}
	l, r := 0, len(A)-1
	for {
		i := (l + r) >> 1
		j := half - i - 2

		var aLeft float64
		if i >= 0 {
			aLeft = float64(A[i])
		} else {
			aLeft = math.Inf(-1)
		}

		var aRight float64
		if (i + 1) < len(A) {
			aRight = float64(A[i+1])
		} else {
			aRight = math.Inf(1)
		}

		var bLeft float64
		if j >= 0 {
			bLeft = float64(B[j])
		} else {
			bLeft = math.Inf(-1)
		}

		var bRight float64
		if (j + 1) < len(B) {
			bRight = float64(B[j+1])
		} else {
			bRight = math.Inf(1)
		}

		if aLeft <= bRight && bLeft <= aRight {
			if total%2 == 0 {
				return float64(math.Max(aLeft, bLeft)+math.Min(aRight, bRight)) / 2
			} else {
				return float64(math.Max(aLeft, bLeft))
			}
		} else if aLeft > bRight {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}

func FindMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	A, B := nums1, nums2
	total := len(nums1) + len(nums2)
	half := (total + 1) / 2

	var Aleft, Aright float64
	var Bleft, Bright float64

	if len(B) < len(A) {
		A, B = B, A
	}

	l, r := 0, len(A)-1
	for {
		i := (l + r) >> 1 // A
		j := half - i - 2 // B

		if i >= 0 {
			Aleft = float64(A[i])
		} else {
			Aleft = math.Inf(-1)
		}

		if (i + 1) < len(A) {
			Aright = float64(A[i+1])
		} else {
			Aright = math.Inf(1)
		}

		if j >= 0 {
			Bleft = float64(B[j])
		} else {
			Bleft = math.Inf(-1)
		}

		if (j + 1) < len(B) {
			Bright = float64(B[j+1])
		} else {
			Bright = math.Inf(1)
		}

		// partition is correct
		if Aleft <= Bright && Bleft <= Aright {
			// odd
			if total%2 == 1 {
				return max(Aleft, Bleft)
			}
			// even
			return (max(Aleft, Bleft) + min(Aright, Bright)) / 2
		} else if Aleft > Bright {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
