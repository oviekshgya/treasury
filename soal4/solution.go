package soal4

import "strconv"

// GO counts decimal digits, ignoring the sign of a negative input.
func GO(num int) map[int]int {
	counts := make(map[int]int)
	for _, digit := range strconv.Itoa(num) {
		if digit != '-' {
			counts[int(digit-'0')]++
		}
	}
	return counts
}
