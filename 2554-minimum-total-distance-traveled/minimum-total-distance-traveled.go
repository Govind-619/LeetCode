func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)

	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})

	n := len(robot)
	m := len(factory)
	const INF int64 = 1e18

	dp := make([][]int64, n+1)
	for i := range dp {
		dp[i] = make([]int64, m+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var solve func(int, int) int64
	solve = func(i, j int) int64 {
		if i == n {
			return 0
		}

		if j == m {
			return INF
		}

		if dp[i][j] != -1 {
			return dp[i][j]
		}

		ans := solve(i, j+1)

		var distance int64 = 0
		pos := factory[j][0]
		limit := factory[j][1]

		for k := 0; k < limit && i+k < n; k++ {
			diff := robot[i+k] - pos
			if diff < 0 {
				diff = -diff
			}

			distance += int64(diff)

			next := solve(i+k+1, j+1)

			if next != INF {
				if distance+next < ans {
					ans = distance + next
				}
			}
		}

		dp[i][j] = ans
		return ans
	}

	return solve(0, 0)
}