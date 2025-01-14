package excercise8

import (
	"math"
	"strings"
)

func stringToIntegerAtoi(s string) int {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	result := 0
	sign := 1
	index := 0

	if s[0] == '-' {
		sign = -1
		index = 1
	} else if s[0] == '+' {
		index = 1
		sign = 1
	}

	for i, ii := range []rune(s) {
		if i < index {
			continue
		}

		digit := int(ii - '0')
		if digit < 0 || digit > 9 {
			break
		}

		if (result > math.MaxInt32/10) || (result == math.MaxInt32/10 && digit > math.MaxInt32%10) {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		result = result*10 + digit
	}

	return result * sign
}
