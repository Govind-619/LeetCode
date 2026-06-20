func maxBuilding(n int, restrictions [][]int) int {
	restrictions = append(restrictions, []int{1, 0})

	restrictions = append(restrictions, []int{n, n - 1})

	sort.Slice(restrictions, func(i, j int) bool {
		return restrictions[i][0] < restrictions[j][0]
	})

	m := len(restrictions)

	for i := 1; i < m; i++ {
		dist := restrictions[i][0] - restrictions[i-1][0]

		restrictions[i][1] = min(
			restrictions[i][1],
			restrictions[i-1][1]+dist,
		)
	}

	for i := m - 2; i >= 0; i-- {
		dist := restrictions[i+1][0] - restrictions[i][0]

		restrictions[i][1] = min(
			restrictions[i][1],
			restrictions[i+1][1]+dist,
		)
	}

	ans := 0

	for i := 1; i < m; i++ {
		h1 := restrictions[i-1][1]
		h2 := restrictions[i][1]

		dist := restrictions[i][0] - restrictions[i-1][0]

		peak := max(h1, h2) +
			(dist-abs(h1-h2))/2

		ans = max(ans, peak)
	}

	return ans
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}