func maxDistance(colors []int) int {
	n := len(colors)
	ans := 0

	for i := 1; i < n; i++ {
		if colors[i] != colors[0] {
			ans = max(ans, i)
		}
		if colors[i] != colors[n-1] {
			ans = max(ans, n-1-i)
		}
	}
	return ans
}