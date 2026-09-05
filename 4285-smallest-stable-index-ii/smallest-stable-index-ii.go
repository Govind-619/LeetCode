func firstStableIndex(nums []int, k int) int {
	n := len(nums)

	prefixMax := make([]int, n)
	suffixMin := make([]int, n)

	prefixMax[0] = nums[0]

	for i := 1; i < n; i++ {
		prefixMax[i] = prefixMax[i-1]
		if nums[i] > prefixMax[i] {
			prefixMax[i] = nums[i]
		}
	}

	suffixMin[n-1] = nums[n-1]

	for i := n - 2; i >= 0; i-- {
		suffixMin[i] = suffixMin[i+1]
		if nums[i] < suffixMin[i] {
			suffixMin[i] = nums[i]
		}
	}

	for i := 0; i < n; i++ {
		instability := prefixMax[i] - suffixMin[i]

		if instability <= k {
			return i
		}
	}

	return -1
}