func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
    const ALPHABET = 26
	const INF = math.MaxInt64 / 4 // prevent overflow

	// 1. Distance matrix: dist[i][j] = min cost to convert i -> j
	dist := make([][]int64, ALPHABET)
	for i := 0; i < ALPHABET; i++ {
		dist[i] = make([]int64, ALPHABET)
		for j := 0; j < ALPHABET; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = INF
			}
		}
	}

	// 2. Initialize graph edges from transformation rules
	for i := 0; i < len(original); i++ {
		u := original[i] - 'a'
		v := changed[i] - 'a'

		// Keep minimum cost if multiple edges exist
		if int64(cost[i]) < dist[u][v] {
			dist[u][v] = int64(cost[i])
		}
	}

	// 3. Floyd–Warshall: all-pairs shortest paths
	for k := 0; k < ALPHABET; k++ {
		for i := 0; i < ALPHABET; i++ {
			for j := 0; j < ALPHABET; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	// 4. Compute total cost for string conversion
	var totalCost int64 = 0

	for i := 0; i < len(source); i++ {
		s := source[i] - 'a'
		t := target[i] - 'a'

		if dist[s][t] == INF {
			return -1 // conversion impossible
		}
		totalCost += dist[s][t]
	}

	return totalCost
}