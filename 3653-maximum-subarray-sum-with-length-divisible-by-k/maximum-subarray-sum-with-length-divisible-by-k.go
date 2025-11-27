func maxSubarraySum(nums []int, k int) int64 {
	ans := int64(math.MinInt64)
	pref := int64(0)
	minPref := make([]int64, k)
	for i := 0; i < k; i++ {
		minPref[i] = math.MaxInt64
	}
	for i, num := range nums {
		pref += int64(num)
		r := (i + 1) % k

		if (i+1)%k == 0 {
			ans = max(ans, pref)
		}

		if minPref[r] != math.MaxInt64 {
			ans = max(ans, pref-minPref[r])
		}

		minPref[r] = min(minPref[r], pref)

	}

	return ans
}