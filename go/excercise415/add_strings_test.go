package excercise415

import "testing"

func TestAddString(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want string
	}{
		{
			s1:   "11",
			s2:   "123",
			want: "134",
		},
		{
			s1:   "456",
			s2:   "77",
			want: "533",
		},
		{
			s1:   "0",
			s2:   "0",
			want: "0",
		},
	}

	for _, tt := range tests {
		got := addStrings(tt.s1, tt.s2)
		if got != tt.want {
			t.Errorf("addStrings() = %v, want %v", got, tt.want)
		}
	}
}
