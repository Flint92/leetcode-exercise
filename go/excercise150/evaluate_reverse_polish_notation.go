package evaluate_reverse_polish_notation

import "strconv"

func evaluateReversePolishNotation(tokens []string) int {
	result := make([]int, 0)
	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			result = append(result, num)
		} else {
			r := result[len(result)-1]
			l := result[len(result)-2]
			result = result[:len(result)-2]

			switch token {
			case "+":
				result = append(result, l+r)
			case "-":
				result = append(result, l-r)
			case "*":
				result = append(result, l*r)
			case "/":
				result = append(result, l/r)
			}
		}
	}
	return result[0]
}
