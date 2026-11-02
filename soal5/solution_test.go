package soal5

import (
	"slices"
	"testing"
)

func TestGO(t *testing.T) {
	for _, tc := range []struct {
		name string
		nums []int
		want int
	}{
		{"even example", []int{1, 2, 3, 4}, 14},
		{"odd example", []int{1, 2, 3, 4, 5}, 39},
		{"single number", []int{5}, 25},
		{"single negative", []int{-5}, 25},
		{"negative product", []int{-2, 3, 0}, -6},
		{"zero product", []int{0, 5}, 0},
		{"empty input", nil, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			original := slices.Clone(tc.nums)
			if got := GO(tc.nums...); got != tc.want {
				t.Errorf("GO(%v) = %d, want %d", tc.nums, got, tc.want)
			}
			if !slices.Equal(tc.nums, original) {
				t.Error("input was modified")
			}
		})
	}
}

func TestGOManyPairs(t *testing.T) {
	nums := make([]int, 1001)
	for i := range nums {
		nums[i] = 1
	}
	if got := GO(nums...); got != 501 {
		t.Errorf("GO with 1001 ones = %d, want 501", got)
	}
}
