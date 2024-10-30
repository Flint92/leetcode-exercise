package word_break

import (
	"testing"
)

func TestWorkBreakDp(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		want     bool
	}{
		{
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			want:     true,
		},
		{
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			want:     true,
		},
		{
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     false,
		},
	}

	for _, tt := range tests {
		if got := wordBreakDp(tt.s, tt.wordDict); got != tt.want {
			t.Errorf("wordBreakDp(%v, %v) = %v, want %v", tt.s, tt.wordDict, got, tt.want)
		}
	}
}
