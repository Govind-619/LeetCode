package main

func computeZ(s string) []int {
	n := len(s)
	z := make([]int, n)
	left, right := 0, 0

	for i := 1; i < n; i++ {
		if i <= right {
			z[i] = min(z[i-left], right-i+1)
		}

		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			left = i
			right = i + z[i]
			z[i]++
		}
	}

	z[0] = n
	return z
}

func generateString(str1 string, str2 string) string {
	n := len(str1)
	m := len(str2)

	// Initialize word with '*'
	word := make([]byte, n+m-1)
	for i := range word {
		word[i] = '*'
	}

	z := computeZ(str2)
	lastMatchStart := -m

	for i := 0; i < n; i++ {
		if str1[i] == 'F' {
			continue
		}

		overlap := max(0, lastMatchStart+m-i)

		if overlap > 0 && z[m-overlap] < overlap {
			return ""
		}

		for j := overlap; j < m; j++ {
			word[i+j] = str2[j]
		}

		lastMatchStart = i
	}

	lastWildcard := make([]int, len(word))
	last := -1

	for i := 0; i < len(word); i++ {
		if word[i] == '*' {
			word[i] = 'a'
			last = i
		}
		lastWildcard[i] = last
	}

	z = computeZ(str2 + string(word))

	for i := 0; i < n; i++ {
		if str1[i] == 'T' || z[m+i] < m {
			continue
		}

		pos := lastWildcard[i+m-1]

		if pos < i {
			return ""
		}

		word[pos] = 'b'
		i = pos
	}

	return string(word)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}