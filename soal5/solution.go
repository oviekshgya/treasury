package soal5

// GO multiplies each adjacent pair in its own goroutine, then sums in the
// calling goroutine. An unpaired last number is multiplied by itself.
func GO(nums ...int) int {
	pairCount := len(nums)/2 + len(nums)%2
	results := make(chan int, pairCount)

	for i := 0; i < len(nums); i += 2 {
		left, right := nums[i], nums[i]
		if i+1 < len(nums) {
			right = nums[i+1]
		}
		go func(left, right int) {
			results <- left * right
		}(left, right)
	}

	total := 0
	for i := 0; i < pairCount; i++ {
		total += <-results
	}
	return total
}
