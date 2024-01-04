package leetcode

import "strings"

func GenerateParanthesis(n int) []string {
	stack := make([]string, 0)
	result := make([]string, 0)
	var backtrack func(int, int)
	backtrack = func(open int, close int) {
		if open == n && close == n {
			result = append(result, strings.Join(stack, ""))
			return
		}

		if open < n {
			stack = append(stack, "(")
			backtrack(open+1, close)
			stack = stack[:len(stack)-1]
		}

		if close < open && close < n {
			stack = append(stack, ")")
			backtrack(open, close+1)
			stack = stack[:len(stack)-1]
		}
	}
	backtrack(0, 0)
	return result
}
