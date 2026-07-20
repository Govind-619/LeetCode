func shiftGrid(grid [][]int, k int) [][]int {

	m := len(grid)
	n := len(grid[0])

	total := m * n

	k %= total

	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			oldIndex := i*n + j

			newIndex := (oldIndex + k) % total

			newRow := newIndex / n
			newCol := newIndex % n

			ans[newRow][newCol] = grid[i][j]
		}
	}

	return ans
}