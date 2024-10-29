package evaluate_reverse_polish_notation

import "testing"

func TestEvaluateReversePolishNotation(t *testing.T) {
	tests := []struct {
		tokens []string
		want   int
	}{
		{
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
	}

	for _, tt := range tests {
		if got := evaluateReversePolishNotation(tt.tokens); got != tt.want {
			t.Errorf("evaluateReversePolishNotation(%v) = %v, want %v", tt.tokens, got, tt.want)
		}
	}
}
