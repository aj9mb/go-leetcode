package leetcode

import "math"

func MinEatingSpeed(piles []int, h int) int {
	l, r := 1, piles[0]
	for _, p := range piles {
		if p > r {
			r = p
		}
	}
	res := r
	for l <= r {
		k := int(math.Floor((float64(l) + float64(r)) / 2))
		var hrs int = 0
		for _, p := range piles {
			hrs += int(math.Ceil(float64(p) / float64(k)))
		}
		if hrs > h {
			l = k + 1
		} else {
			if res > k {
				res = k
			}
			r = k - 1
		}
	}
	return res
}
