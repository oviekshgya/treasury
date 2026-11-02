package soal3

import (
	"slices"
	"testing"
)

func TestGO(t *testing.T) {
	for _, tc := range []struct {
		name       string
		a, b, want []int
	}{
		{"single match", []int{1, 2}, []int{1, 3}, []int{1}},
		{"duplicate example", []int{1, 2, 2}, []int{1, 2, 4}, []int{2, 1}},
		{"duplicates in both inputs", []int{-2, 0, 3, 3}, []int{3, -2, 3}, []int{3, -2}},
		{"zero and descending order", []int{-1, 0, 4, 2}, []int{2, 4, 0, -1}, []int{4, 2, 0, -1}},
		{"no matches", []int{1}, []int{2}, []int{}},
		{"empty first input", nil, []int{1}, []int{}},
		{"empty second input", []int{1}, nil, []int{}},
		{"empty inputs", nil, nil, []int{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			originalA, originalB := slices.Clone(tc.a), slices.Clone(tc.b)
			if got := GO(tc.a, tc.b); !slices.Equal(got, tc.want) {
				t.Errorf("GO(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
			if !slices.Equal(tc.a, originalA) || !slices.Equal(tc.b, originalB) {
				t.Error("input was modified")
			}
		})
	}
}
