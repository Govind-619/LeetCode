func numberOfSubmatrices(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var line = make([]int, n+1)
	for j := 0; j <= n; j++ {
		line[j] = 1 << 23
	}
	mask, unmask := 1<<25, (1<<25)-1

	var ret = 0
	for i := 0; i < m; i++ {
		var lastInc = 1 << 23
		for j := 1; j <= n; j++ {
			xMask := (line[j] & mask) | (line[j-1] & mask)
			var inc = 0
			if grid[i][j-1] == 'X' {
				inc = 1
				xMask = mask
			} else if grid[i][j-1] == 'Y' {
				inc = -1
			}
			newVal := (line[j] & unmask) + (line[j-1] & unmask) - lastInc + inc - (1 << 23)
			if xMask > 0 && newVal == 0 {
				ret++
			}
			lastInc = line[j] & unmask
			line[j] = (newVal + (1 << 23)) | xMask
		}
	}
	return ret
}