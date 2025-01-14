package excercise8

import (
	"math"
	"unicode"
)

func stringToIntegerAtoi(s string) int {
	stack := make([]int, 0, len(s))
	sign := 1

	for _, ch := range []rune(s) {
		if unicode.IsSpace(ch) {
			continue
		}
		if ch == '-' {
			if len(stack) == 0 && sign == 1 {
				sign = -1
				continue
			} else {
				break
			}
		}
		if unicode.IsLetter(ch) {
			break
		}
		if unicode.IsDigit(ch) {
			stack = append(stack, int(ch-'0'))
		}
	}

	result := 0
	for _, ii := range stack {
		if (result > math.MaxInt32/10) || (result == math.MaxInt32/10 && ii > math.MaxInt32%10) {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		result = result*10 + ii
	}

	return result * sign
}
