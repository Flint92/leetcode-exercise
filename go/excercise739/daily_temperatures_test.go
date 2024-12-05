package daily_temperatures

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		temperatures []int
		want         []int
	}{
		{
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:         []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			temperatures: []int{30, 40, 50, 60},
			want:         []int{1, 1, 1, 0},
		},
		{
			temperatures: []int{30, 60, 90},
			want:         []int{1, 1, 0},
		},
	}

	for _, tt := range tests {
		if got := dailyTemperatures(tt.temperatures); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("dailyTemperatures(%v) = %v, want %v", tt.temperatures, got, tt.want)
		}
	}
}
