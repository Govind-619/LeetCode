func maxJumps(arr []int, d int) int {
	n := len(arr)
	dp := make([]int, n)

	var dfs func(int) int

	dfs = func(i int) int {
		if dp[i] != 0 {
			return dp[i]
		}

		best := 1

		for nxt := i + 1; nxt <= min(n-1, i+d); nxt++ {
			if arr[nxt] >= arr[i] {
				break
			}

			best = max(best, 1+dfs(nxt))
		}

		for nxt := i - 1; nxt >= max(0, i-d); nxt-- {
			if arr[nxt] >= arr[i] {
				break
			}

			best = max(best, 1+dfs(nxt))
		}

		dp[i] = best
		return dp[i]
	}

	ans := 1

	for i := 0; i < n; i++ {
		ans = max(ans, dfs(i))
	}

	return ans
}