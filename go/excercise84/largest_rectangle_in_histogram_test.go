package largest_rectangle_in_histogram

import (
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{
			heights: []int{2, 1, 5, 6, 2, 3},
			want:    10,
		},
	}

	for _, tt := range tests {
		if got := largestRectangleArea(tt.heights); got != tt.want {
			t.Errorf("largestRectangleArea(%v) = %v, want %v", tt.heights, got, tt.want)
		}
	}
}
