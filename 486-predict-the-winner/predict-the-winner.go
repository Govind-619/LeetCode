func predictTheWinner(nums []int) bool {
	n := len(nums)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = nums[i]
	}

	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1

			takeLeft := nums[i] - dp[i+1][j]

			takeRight := nums[j] - dp[i][j-1]

			if takeLeft > takeRight {
				dp[i][j] = takeLeft
			} else {
				dp[i][j] = takeRight
			}
		}
	}

	return dp[0][n-1] >= 0
}