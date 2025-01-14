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
	digitCount := 0
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		result += top * int(math.Pow10(digitCount))
		digitCount++
		stack = stack[:len(stack)-1]
	}

	return result * sign
}
