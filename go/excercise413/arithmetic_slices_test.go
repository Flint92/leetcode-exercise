package arithmetic_slices

import (
	"testing"
)

func TestNumberOfArithmeticSlices(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 2, 3, 4}, want: 3},
		{nums: []int{1}, want: 0},
	}

	for _, test := range tests {
		if got := numberOfArithmeticSlices(test.nums); got != test.want {
			t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, test.want)
		}
	}
}
