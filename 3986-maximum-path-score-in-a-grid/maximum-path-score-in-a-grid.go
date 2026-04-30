func maxPathScore(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])
	const NEG = -1000000000

	prev := make([][]int, n)
	for j := 0; j < n; j++ {
		prev[j] = make([]int, k+1)
		for c := 0; c <= k; c++ {
			prev[j][c] = NEG
		}
	}

	for i := 0; i < m; i++ {
		curr := make([][]int, n)
		for j := 0; j < n; j++ {
			curr[j] = make([]int, k+1)
			for c := 0; c <= k; c++ {
				curr[j][c] = NEG
			}
		}

		for j := 0; j < n; j++ {
			gain := grid[i][j]
			need := 0
			if gain > 0 {
				need = 1
			}

			limit := k
			if i+j < limit {
				limit = i + j
			}

			if i == 0 && j == 0 {
				curr[0][0] = 0
				continue
			}

			for c := need; c <= limit; c++ {
				best := NEG

				if i > 0 && prev[j][c-need] != NEG {
					val := prev[j][c-need] + gain
					if val > best {
						best = val
					}
				}

				if j > 0 && curr[j-1][c-need] != NEG {
					val := curr[j-1][c-need] + gain
					if val > best {
						best = val
					}
				}

				curr[j][c] = best
			}
		}

		prev = curr
	}

	ans := NEG
	for c := 0; c <= k; c++ {
		if prev[n-1][c] > ans {
			ans = prev[n-1][c]
		}
	}

	if ans < 0 {
		return -1
	}
	return ans
}