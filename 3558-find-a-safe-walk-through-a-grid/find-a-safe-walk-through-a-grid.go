func findSafeWalk(grid [][]int, health int) bool {
	m := len(grid)
	n := len(grid[0])

	const INF = int(1e9)

	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = INF
		}
	}

	type Pair struct {
		x, y int
	}

	deque := make([]Pair, 0)
	head := 0

	dist[0][0] = grid[0][0]
	deque = append(deque, Pair{0, 0})

	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	for head < len(deque) {
		cur := deque[head]
		head++

		for k := 0; k < 4; k++ {
			nx := cur.x + dx[k]
			ny := cur.y + dy[k]

			if nx < 0 || ny < 0 || nx >= m || ny >= n {
				continue
			}

			newCost := dist[cur.x][cur.y] + grid[nx][ny]

			if newCost < dist[nx][ny] {
				dist[nx][ny] = newCost

				if grid[nx][ny] == 0 {
					deque = append([]Pair{{nx, ny}}, deque[head:]...)
					head = 0
				} else {
					deque = append(deque, Pair{nx, ny})
				}
			}
		}
	}

	return dist[m-1][n-1] < health
}