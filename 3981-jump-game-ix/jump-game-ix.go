func maxValue(nums []int) []int {
	n := len(nums)

	suffixMin := make([]int, n+1)
	suffixMin[n] = int(^uint(0) >> 1)

	for i := n - 1; i >= 0; i-- {
		if nums[i] < suffixMin[i+1] {
			suffixMin[i] = nums[i]
		} else {
			suffixMin[i] = suffixMin[i+1]
		}
	}

	ans := make([]int, n)
	l := 0

	for l < n {
		r := l
		componentMax := nums[l]

		for r+1 < n && componentMax > suffixMin[r+1] {
			r++
			if nums[r] > componentMax {
				componentMax = nums[r]
			}
		}

		for i := l; i <= r; i++ {
			ans[i] = componentMax
		}

		l = r + 1
	}

	return ans
}