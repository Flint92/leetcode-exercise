package excercise415

import (
	"strconv"
	"strings"
)

func addStrings(s1, s2 string) string {
	i := len(s1) - 1
	j := len(s2) - 1
	carry := 0

	result := make([]string, 0, maxInt(i, j)+1)
	for i >= 0 || j >= 0 || carry != 0 {
		x, y := 0, 0

		if i >= 0 {
			x = int(s1[i] - '0')
			i--
		}

		if j >= 0 {
			y = int(s2[j] - '0')
			j--
		}

		sum := x + y + carry
		result = append([]string{strconv.Itoa(sum % 10)}, result...)
		carry = sum / 10
	}

	return strings.Join(result, "")
}

func maxInt(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
