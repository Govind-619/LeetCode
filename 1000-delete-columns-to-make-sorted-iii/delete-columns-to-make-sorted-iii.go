// Complexity:
// Time O(NM^2) and Space O(M) where N is the length of strs and M
// is the length of each word.
func minDeletionSize(strs []string) int {
	m := len(strs[0])
	result := m - 1

	// dp[i]:= the minimum deletions for strs[:][i:] if we keep the
	// column i.
	dp := make([]int, m)
	for first := m - 2; first >= 0; first-- {
		dp[first] = m - first - 1
		for next := first + 1; next-first-1 < dp[first]; next++ {
			if isValidNext(strs, first, next) {
				dp[first] = min(dp[first], next-first-1+dp[next])
			}
		}

		result = min(result, first+dp[first])
	}

	return result
}

// isValidNext returns whether it satisfies the lexicographic order
// for the first two columns.
func isValidNext(strs []string, first, next int) bool {
	for row := range strs {
		if strs[row][first] > strs[row][next] {
			return false
		}
	}
	return true
}