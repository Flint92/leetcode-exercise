package maximum_product_of_word_lengths

func maximumProductOfWordLengths(words []string) int {
	n := len(words)
	bits := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		for _, char := range words[i] {
			bits[i] |= 1 << (char - 'a')
		}

		for j := 0; j < i; j++ {
			if bits[i]&bits[j] == 0 {
				prod := len(words[i]) * len(words[j])
				if prod > ans {
					ans = prod
				}
			}
		}
	}
	return ans
}
