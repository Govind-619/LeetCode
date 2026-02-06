func minRemoval(nums []int, k int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	sort.Ints(nums)

	maxLen := 1
	l := 0

	for r := 0; r < n; r++ {
		for nums[r] > nums[l]*k {
			l++
		}
		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}
	}

	return n - maxLen
}
