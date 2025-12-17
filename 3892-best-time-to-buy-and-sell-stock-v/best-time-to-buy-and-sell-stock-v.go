func maximumProfit(prices []int, K int) int64 {
	n := len(prices)
	maxProfit := 0

	dp := make([]int, n)

	for k := 1; k <= K; k++ {
		maxT1 := -prices[0]
		maxT2 := prices[0]

		prev := dp[k-1]

		for i := k; i < n; i++ {
			current := dp[i]

			dp[i] = max(maxT1+prices[i], maxT2-prices[i], dp[i-1])
			maxProfit = max(maxProfit, dp[i])

			maxT1 = max(maxT1, prev-prices[i])
			maxT2 = max(maxT2, prev+prices[i])

			prev = current
		}
	}

	return int64(maxProfit)
}
