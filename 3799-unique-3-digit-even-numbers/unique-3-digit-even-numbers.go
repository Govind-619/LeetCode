func totalNumbers(digits []int) int {
	freq := make([]int, 10)

	for _, digit := range digits {
		freq[digit]++
	}

	answer := 0

	for first := 1; first <= 9; first++ {
		for second := 0; second <= 9; second++ {
			for third := 0; third <= 8; third += 2 {
				if freq[first] == 0 || freq[second] == 0 || freq[third] == 0 {
					continue
				}

				if first == second && second == third && freq[first] < 3 {
					continue
				}

				if first == second && freq[first] < 2 {
					continue
				}

				if first == third && freq[first] < 2 {
					continue
				}

				if second == third && freq[second] < 2 {
					continue
				}

				answer++
			}
		}
	}

	return answer
}