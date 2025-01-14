package excercise8

import "testing"

func TestStringToIntegerAtoi(t *testing.T) {
	tests := []struct {
		Input  string
		Output int
	}{
		{
			Input:  "42",
			Output: 42,
		},
		{
			Input:  "  -42 ",
			Output: -42,
		},
		{
			Input:  "1337c0d3",
			Output: 1337,
		},
		{
			Input:  "0-1",
			Output: 0,
		},
		{
			Input:  "words and 987",
			Output: 0,
		},
	}

	for _, test := range tests {
		if got := stringToIntegerAtoi(test.Input); got != test.Output {
			t.Errorf("StringToIntegerAtoi(%q) = %d, want %d", test.Input, got, test.Output)
		}
	}
}
