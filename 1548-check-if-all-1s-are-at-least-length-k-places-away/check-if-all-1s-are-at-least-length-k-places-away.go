func kLengthApart(nums []int, k int) bool {
	last := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if last != -1 && i-last-1 < k {
				return false
			} else {
				last = i
			}
		}
	}
	return true
}