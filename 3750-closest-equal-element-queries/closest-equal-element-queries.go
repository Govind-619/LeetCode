func solveQueries(nums []int, queries []int) []int {
	distanceLeft := make([]int, len(nums))
    distanceRight := make([]int, len(nums))
	lastSeenLeft := make(map[int]int, len(nums))
    lastSeenRight := make(map[int]int, len(nums))

	for i, j := 0,(len(nums)*2)-1; i < len(nums)*2; i, j = i+1, j-1 {
		if v, ok := lastSeenLeft[nums[i%len(nums)]]; ok {
            // If its not same element
			if v != i%len(nums) {
				distanceLeft[i%len(nums)] = i - v
				lastSeenLeft[nums[i%len(nums)]] = i
			}
		} else {
			distanceLeft[i%len(nums)] = -1
			lastSeenLeft[nums[i%len(nums)]] = i
		}
        if v, ok := lastSeenRight[nums[j%len(nums)]]; ok {
            // If its not same element
			if v != j+len(nums) {
				distanceRight[j%len(nums)] = v - j
				lastSeenRight[nums[j%len(nums)]] = j
			}
		} else {
			distanceRight[j%len(nums)] = -1
			lastSeenRight[nums[j%len(nums)]] = j
		}
	}
	
    res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = min(distanceLeft[q], distanceRight[q])
	}
	return res
}