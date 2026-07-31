func minimumPushes(word string) int {
	freq := make([]int, 26)

	for _, ch := range word {
		freq[ch-'a']++
	}

	sort.Slice(freq, func(i, j int) bool {
		return freq[i] > freq[j]
	})

	ans := 0

	for i := 0; i < 26; i++ {
		if freq[i] == 0 {
			break
		}

		pushes := (i / 8) + 1

		ans += freq[i] * pushes
	}

	return ans
}