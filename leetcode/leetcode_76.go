package leetcode

// func MinWindow(s string, t string) string {
// 	op := ""
// 	sArr := make(map[rune]int)
// 	tMap := make(map[rune]int)
// 	for _, tStr := range t {
// 		k := tStr
// 		tMap[k]++
// 	}
// 	l, r := 0, 0
// 	if len(s) < len(t) {
// 		r = 1
// 	} else {
// 		r = len(t)
// 	}

// 	for i := 0; i < r; i++ {
// 		sArr[rune(s[i])]++
// 	}
// 	for l < r {
// 		if matchStrWithTarget(tMap, sArr) {
// 			tmpOp := s[l:r]
// 			if op == "" || len(op) > len(tmpOp) {
// 				op = tmpOp
// 			}
// 			sArr[rune(s[l])]--
// 			l++
// 		} else {
// 			if r+1 > len(s) {
// 				break
// 			}
// 			sArr[rune(s[r])]++
// 			r++
// 		}
// 	}
// 	return op
// }

// func matchStrWithTarget(target map[rune]int, sArr map[rune]int) bool {
// 	res := false
// 	for k, v := range target {
// 		if sArr[k] >= v {
// 			res = true
// 		} else {
// 			return false
// 		}
// 	}
// 	return res
// }

func MinWindow(s string, t string) string {
	op := ""
	if s == "" || t == "" {
		return op
	}
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)
	for _, v := range t {
		tMap[v]++
	}
	l, r, hCount, nCount := 0, 0, 0, len(tMap)
	for l <= r {
		if hCount == nCount {
			if op == "" || len(op) > (r-l) {
				op = s[l:r]
			}
			ln := rune(s[l])
			sMap[ln]--
			if sMap[ln] < tMap[ln] {
				hCount--
			}
			l++
		} else {
			if r >= len(s) {
				break
			}
			rn := rune(s[r])
			sMap[rn]++
			if tMap[rn] == sMap[rn] {
				hCount++
			}
			r++
		}
	}
	return op
}
