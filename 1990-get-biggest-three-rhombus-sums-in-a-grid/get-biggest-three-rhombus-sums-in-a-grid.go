func getBiggestThree(grid [][]int) []int {
	m := len(grid)
	n := len(grid[0])

	set := map[int]bool{}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {

			set[grid[r][c]] = true

			for k := 1; ; k++ {

				if r+2*k >= m || c-k < 0 || c+k >= n {
					break
				}

				sum := 0

				x, y := r, c
				for i := 0; i < k; i++ {
					sum += grid[x+i][y+i]
				}

				x, y = r+k, c+k
				for i := 0; i < k; i++ {
					sum += grid[x+i][y-i]
				}

				x, y = r+2*k, c
				for i := 0; i < k; i++ {
					sum += grid[x-i][y-i]
				}

				x, y = r+k, c-k
				for i := 0; i < k; i++ {
					sum += grid[x-i][y+i]
				}

				set[sum] = true
			}
		}
	}

	arr := []int{}
	for v := range set {
		arr = append(arr, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	if len(arr) > 3 {
		arr = arr[:3]
	}

	return arr
}