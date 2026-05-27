func numberOfSpecialChars(word string) int {
	lastLower := make([]int, 26)
	firstUpper := make([]int, 26)

	for i := 0; i < 26; i++ {
		lastLower[i] = -1
		firstUpper[i] = -1
	}

	invalid := make(map[int]bool)

	for i, ch := range word {
		if ch >= 'a' && ch <= 'z' {
			idx := ch - 'a'

			lastLower[idx] = i

			if firstUpper[idx] != -1 {
				invalid[int(idx)] = true
			}

		} else {
			idx := ch - 'A'

			if firstUpper[idx] == -1 {
				firstUpper[idx] = i
			}
		}
	}

	specialCount := 0

	for i := 0; i < 26; i++ {
		if lastLower[i] != -1 &&
			firstUpper[i] != -1 &&
			!invalid[i] {

			specialCount++
		}
	}

	return specialCount
}