func rotateTheBox(boxGrid [][]byte) [][]byte {
	m, n := len(boxGrid), len(boxGrid[0])

	// gravity to the right
	for i := 0; i < m; i++ {
		write := n - 1
		for j := n - 1; j >= 0; j-- {
			if boxGrid[i][j] == '*' {
				write = j - 1
			} else if boxGrid[i][j] == '#' {
				boxGrid[i][j] = '.'
				boxGrid[i][write] = '#'
				write--
			}
		}
	}

	//rotate 90 degree to the right
	res := make([][]byte, n)
	for i := range res {
		res[i] = make([]byte, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][m-1-i] = boxGrid[i][j]
		}
	}
	return res
}