package combination_sum_iii

import (
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	tests := []struct {
		k    int
		n    int
		want [][]int
	}{
		{
			k:    3,
			n:    7,
			want: [][]int{{1, 2, 4}},
		},
		{
			k:    3,
			n:    9,
			want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		},
		{
			k:    4,
			n:    1,
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		if got := combinationSum3(tt.k, tt.n); !areSlicesEqual(got, tt.want) {
			t.Errorf("combinationSum3(%v, %v) = %v, want %v", tt.k, tt.n, got, tt.want)
		}
	}

}

func areSlicesEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
