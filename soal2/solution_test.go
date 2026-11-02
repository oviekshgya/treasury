package soal2

import (
	"errors"
	"slices"
	"strconv"
	"testing"
)

func TestGO(t *testing.T) {
	for _, tc := range []struct {
		name      string
		exclude   []int
		nums      []int
		want      int
		wantError string
	}{
		{"all digits excluded", []int{0, 1, 0}, []int{1000, 1010}, 0, "not found"},
		{"filtered example", []int{0, 1, 0}, []int{1259, 2601}, 62952, ""},
		{"reverse combined digits", nil, []int{12, 34}, 4321, ""},
		{"leading zeros in result", nil, []int{120}, 21, ""},
		{"leading zeros in combined number", nil, []int{0, 0, 12}, 21, ""},
		{"zero between numbers", nil, []int{12, 0, 34}, 43021, ""},
		{"internal zeros preserved", nil, []int{102, 304}, 403201, ""},
		{"zero remains", []int{1}, []int{100}, 0, ""},
		{"zero input", nil, []int{0}, 0, ""},
		{"all zero inputs", nil, []int{0, 0, 0}, 0, ""},
		{"all zero inputs excluded", []int{0}, []int{0, 0, 0}, 0, "not found"},
		{"empty input", nil, nil, 0, "not found"},
		{"negative input", []int{1}, []int{-123}, 32, ""},
		{"invalid exclusion values", []int{-1, 10, 12}, []int{12}, 21, ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			originalExclude, originalNums := slices.Clone(tc.exclude), slices.Clone(tc.nums)
			got, err := GO(tc.exclude, tc.nums...)
			if got != tc.want {
				t.Errorf("result = %d, want %d", got, tc.want)
			}
			if tc.wantError != "" {
				if !errors.Is(err, ErrNotFound) || err.Error() != tc.wantError {
					t.Errorf("error = %v, want %q", err, tc.wantError)
				}
			} else if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !slices.Equal(tc.exclude, originalExclude) || !slices.Equal(tc.nums, originalNums) {
				t.Error("input was modified")
			}
		})
	}
}

func TestGOOverflow(t *testing.T) {
	maxInt := int(^uint(0) >> 1)
	got, err := GO(nil, maxInt, maxInt)
	if !errors.Is(err, strconv.ErrRange) {
		t.Fatalf("expected integer overflow, got %v", err)
	}
	if got != 0 {
		t.Errorf("result on overflow = %d, want 0", got)
	}
}
