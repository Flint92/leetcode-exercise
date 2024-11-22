package maximum_product_of_word_lengths

import "testing"

func TestMaximumProductOfWordLengths(t *testing.T) {
	tests := []struct {
		words []string
		want  int
	}{
		{
			words: []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"},
			want:  16,
		},
		{
			words: []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"},
			want:  4,
		},
		{
			words: []string{"a", "aa", "aaa", "aaaa"},
			want:  0,
		},
	}

	for _, tt := range tests {
		if got := maximumProductOfWordLengths(tt.words); got != tt.want {
			t.Errorf("maximumProductOfWordLengths(%v) = %v, want %v", tt.words, got, tt.want)
		}
	}
}
