package soal4

import (
	"maps"
	"strconv"
	"testing"
)

func TestGO(t *testing.T) {
	for _, tc := range []struct {
		name string
		num  int
		want map[int]int
	}{
		{"example", 1223334, map[int]int{1: 1, 2: 2, 3: 3, 4: 1}},
		{"zero", 0, map[int]int{0: 1}},
		{"negative", -1002, map[int]int{0: 2, 1: 1, 2: 1}},
		{"single digit", 7, map[int]int{7: 1}},
		{"repeated digit", 9999, map[int]int{9: 4}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := GO(tc.num); !maps.Equal(got, tc.want) {
				t.Errorf("GO(%d) = %v, want %v", tc.num, got, tc.want)
			}
		})
	}
}

func TestGOMinInt(t *testing.T) {
	minInt := -int(^uint(0)>>1) - 1
	want := map[int]int{1: 1, 2: 1, 3: 1, 4: 3, 6: 1, 7: 1, 8: 2}
	if strconv.IntSize == 64 {
		want = map[int]int{0: 2, 2: 3, 3: 3, 4: 1, 5: 2, 6: 1, 7: 3, 8: 3, 9: 1}
	}
	if got := GO(minInt); !maps.Equal(got, want) {
		t.Errorf("GO(%d) = %v, want %v", minInt, got, want)
	}
}
