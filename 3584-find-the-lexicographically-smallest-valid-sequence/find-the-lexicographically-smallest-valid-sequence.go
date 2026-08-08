func validSequence(word1 string, word2 string) []int {
	n := len(word1)
	m := len(word2)

	last := make([]int, m)

	for i := 0; i < m; i++ {
		last[i] = -1
	}

	i := n - 1
	j := m - 1

	for i >= 0 && j >= 0 {
		if word1[i] == word2[j] {
			last[j] = i
			j--
		}

		i--
	}

	ans := make([]int, 0, m)
	canSkip := true
	j = 0

	for i := 0; i < n && j < m; i++ {
		if word1[i] == word2[j] {
			ans = append(ans, i)
			j++
		} else if canSkip && (j == m-1 || i < last[j+1]) {
			canSkip = false
			ans = append(ans, i)
			j++
		}
	}

	if j == m {
		return ans
	}

	return []int{}
}