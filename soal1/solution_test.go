package soal1

import (
	"slices"
	"testing"
)

func TestGO(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []int
		want  int
	}{
		{"mixed signs", []int{-1, 1}, 0},
		{"negative example", []int{-1, -7, -5}, -2},
		{"positive example", []int{1, 2, 1, 6}, 3},
		{"missing positive one", []int{0, 2, 3}, 1},
		{"missing negative one", []int{0, -2, -3}, -1},
		{"consecutive positives", []int{3, 1, 2}, 4},
		{"consecutive negatives", []int{-3, -1, -2, -1}, -4},
		{"mixed signs with zero", []int{0, 1, -2}, 0},
		{"only zeros", []int{0, 0}, 0},
		{"nil input", nil, 0},
		{"empty input", []int{}, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			original := slices.Clone(tc.input)
			if got := GO(tc.input); got != tc.want {
				t.Errorf("GO(%v) = %d, want %d", tc.input, got, tc.want)
			}
			if !slices.Equal(tc.input, original) {
				t.Error("input was modified")
			}
		})
	}
}
