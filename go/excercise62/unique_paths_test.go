package excercise62

import (
	"log"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m    int
		n    int
		want int
	}{
		{
			m:    3,
			n:    7,
			want: 28,
		},
		{
			m:    3,
			n:    2,
			want: 3,
		},
	}

	for _, tt := range tests {
		if got := uniquePaths(tt.m, tt.n); got != tt.want {
			log.Fatalf("uniquePaths(%d, %d) = %d, want %d", tt.m, tt.n, got, tt.want)
		}
	}
}
