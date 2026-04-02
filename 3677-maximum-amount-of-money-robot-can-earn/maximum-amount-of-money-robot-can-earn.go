func maximumAmount(g [][]int) int {
	h := len(g)
	w := len(g[0])

	// dp[x][k] where:
	// k = 0,1,2 (states as in original code)
	dp := make([][3]int, w)

	// Initialize with very small value
	for i := 0; i < w; i++ {
		dp[i][0] = math.MinInt32
		dp[i][1] = math.MinInt32
		dp[i][2] = math.MinInt32
	}

	// Equivalent to: fill_n(dp[0], 3, 0)
	dp[0][0], dp[0][1], dp[0][2] = 0, 0, 0

	for y := 0; y < h; y++ {
		// x = 0 (first column)
		{
			top0, top1, top2 := dp[0][0], dp[0][1], dp[0][2]
			v := g[y][0]
			v0 := max(v, 0)

			dp[0][0] = v + top0
			dp[0][1] = max(v+top1, v0+top0)
			dp[0][2] = max(v+top2, v0+top1)
		}

		// rest of the row
		for x := 1; x < w; x++ {
			top0, top1, top2 := dp[x][0], dp[x][1], dp[x][2]
			v := g[y][x]
			v0 := max(v, 0)

			dp[x][0] = max(v+dp[x-1][0], v+top0)

			dp[x][1] = max(
				max(v+dp[x-1][1], v+top1),
				max(v0+dp[x-1][0], v0+top0),
			)

			dp[x][2] = max(
				max(v+dp[x-1][2], v+top2),
				max(v0+dp[x-1][1], v0+top1),
			)
		}
	}

	return dp[w-1][2]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}