
const mod int = 1_000_000_007

func numberOfPaths(grid [][]int, k int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, k)
		}
	}

	dp[0][0][grid[0][0]%k] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for prev := 0; prev < k; prev++ {
				curr := (prev + grid[i][j]) % k
				if i-1 >= 0 {
					dp[i][j][curr] = (dp[i][j][curr] + dp[i-1][j][prev]) % mod
				}
				if j-1 >= 0 {
					dp[i][j][curr] = (dp[i][j][curr] + dp[i][j-1][prev]) % mod
				}
			}
		}
	}

	return dp[m-1][n-1][0]
}