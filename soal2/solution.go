package soal2

import (
	"errors"
	"strconv"
	"strings"
)

// ErrNotFound means no digits remain after applying the exclusion list.
var ErrNotFound = errors.New("not found")

// GO concatenates numbers, reverses the combined number, and removes excluded digits.
// Minus signs and exclusion values outside 0-9 are ignored.
// No remaining digits returns ErrNotFound; overflow returns strconv.ErrRange
// wrapped in a *strconv.NumError. Both errors return a zero result.
func GO(exclude []int, nums ...int) (int, error) {
	var excluded [10]bool
	for _, digit := range exclude {
		if digit >= 0 && digit <= 9 {
			excluded[digit] = true
		}
	}

	// A combined number has no leading zeros, except when the number itself is 0.
	for len(nums) > 1 && nums[0] == 0 {
		nums = nums[1:]
	}

	var result strings.Builder
	for i := len(nums) - 1; i >= 0; i-- {
		digits := strconv.Itoa(nums[i])
		for j := len(digits) - 1; j >= 0; j-- {
			if digits[j] == '-' {
				continue
			}
			if !excluded[digits[j]-'0'] {
				result.WriteByte(digits[j])
			}
		}
	}

	if result.Len() == 0 {
		return 0, ErrNotFound
	}

	number, err := strconv.Atoi(result.String())
	if err != nil {
		return 0, err
	}
	return number, nil
}
