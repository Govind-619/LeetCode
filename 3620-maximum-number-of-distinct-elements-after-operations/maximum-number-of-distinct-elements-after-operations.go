func maxDistinctElements(nums []int, k int) int {
	ans := 1
	sort.Ints(nums)
	last := nums[0] - k
	for i := 1; i < len(nums); i++ {
		goal := max(last+1, nums[i]-k)
		if goal > nums[i]+k {
			continue
		}
		last = goal
		ans++
	}
	return ans
}