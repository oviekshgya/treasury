package soal3

import "sort"

// GO returns the distinct numbers present in both slices, in descending order.
func GO(numsA, numsB []int) []int {
	present := make(map[int]struct{}, len(numsB))
	for _, number := range numsB {
		present[number] = struct{}{}
	}

	result := make([]int, 0)
	for _, number := range numsA {
		if _, exists := present[number]; exists {
			result = append(result, number)
			delete(present, number)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(result)))
	return result
}
