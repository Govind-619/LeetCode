func maxProfit(prices []int, strategy []int, k int) int64 {
	// time: O(n), space: O(1)

	profit := int64(0)
	for i := range prices {
		profit += int64(prices[i] * strategy[i])
	}
	maxProfit := profit
	k2 := k / 2
	for i := range k {
		profit -= int64(prices[i] * strategy[i])
		if i >= k2 {
			profit += int64(prices[i])
		}
	}
	maxProfit = max(maxProfit, profit)
	for i, l := k, len(prices); i < l; i++ {
		profit += int64(prices[i] - prices[i]*strategy[i] +
			prices[i-k]*strategy[i-k] - prices[i-k2])
		maxProfit = max(maxProfit, profit)
	}
	return maxProfit
}