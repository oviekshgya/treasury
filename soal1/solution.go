package soal1

// GO returns the first missing positive or negative integer nearest zero.
// A mixture of positive and negative numbers, or no nonzero numbers, returns 0.
func GO(arr []int) int {
	seen := make(map[int]struct{}, len(arr))
	hasPositive, hasNegative := false, false
	for _, number := range arr {
		switch {
		case number > 0:
			hasPositive = true
		case number < 0:
			hasNegative = true
		default:
			continue
		}
		if hasPositive && hasNegative {
			return 0
		}
		seen[number] = struct{}{}
	}

	if !hasPositive && !hasNegative {
		return 0
	}

	step := 1
	if hasNegative {
		step = -1
	}
	for candidate := step; ; candidate += step {
		if _, exists := seen[candidate]; !exists {
			return candidate
		}
	}
}
