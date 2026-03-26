

func canPartitionGrid(grid [][]int) bool {
	if check(grid) {
		return true
	}

	n, m := len(grid), len(grid[0])

	// transpose matrix
	t := make([][]int, m)
	for i := range t {
		t[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			t[j][i] = grid[i][j]
		}
	}

	return check(t)
}

func check(g [][]int) bool {
	n, m := len(g), len(g[0])

	freqTop := make([]int, 100001)
	freqBottom := make([]int, 100001)

	var total int64 = 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			val := g[i][j]
			total += int64(val)
			freqBottom[val]++
		}
	}

	var topSum int64 = 0
	bottomSum := total

	for i := 0; i < n-1; i++ {
		for j := 0; j < m; j++ {
			val := g[i][j]

			topSum += int64(val)
			bottomSum -= int64(val)

			freqTop[val]++
			freqBottom[val]--
		}

		if topSum == bottomSum {
			return true
		}

		if topSum > bottomSum {
			diff := topSum - bottomSum

			if diff <= 100000 && canRemove(g, freqTop, diff, i+1, m, true) {
				return true
			}
		} else {
			diff := bottomSum - topSum

			if diff <= 100000 && canRemove(g, freqBottom, diff, n-i-1, m, false) {
				return true
			}
		}
	}

	return false
}

func canRemove(g [][]int, freq []int, diff int64, h, w int, isTop bool) bool {

	if h > 1 && w > 1 {
		return freq[diff] > 0
	}

	if w == 1 {
		if isTop {
			return int64(g[0][0]) == diff || int64(g[h-1][0]) == diff
		} else {
			n := len(g)
			return int64(g[n-h][0]) == diff || int64(g[n-1][0]) == diff
		}
	}

	if h == 1 {
		if isTop {
			return int64(g[0][0]) == diff || int64(g[0][w-1]) == diff
		} else {
			n := len(g)
			return int64(g[n-1][0]) == diff || int64(g[n-1][w-1]) == diff
		}
	}

	return false
}