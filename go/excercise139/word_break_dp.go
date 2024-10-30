package word_break

import "slices"

func wordBreakDp(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && slices.Contains(wordDict, s[j:i]) {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}
