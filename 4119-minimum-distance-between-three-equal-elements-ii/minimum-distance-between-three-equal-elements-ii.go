// minimumDistance calculates the minimum distance.
func minimumDistance(nums []int) int {
	// Map to store the most recent index (idx_m+1) of a number
	lastSeen := make(map[int]int)
	// Map to store the second most recent index (idx_m) of a number
	secondLastSeen := make(map[int]int)

	// Initialize minimum distance to a very large value (MaxInt32 is safe here)
	minDistance := math.MaxInt32

	for i, value := range nums {
		// Check if we have seen this number at least twice before
		if firstIndex, ok := secondLastSeen[value]; ok {
			// Case 1: Third (or more) occurrence found.
			// i is the third index (k).
			// firstIndex is the first index of this triplet (i).

			// Calculate distance: 2 * (k - i)
			currentDist := 2 * (i - firstIndex)

			// Update global minimum
			if currentDist < minDistance {
				minDistance = currentDist
			}

			// Slide the window:
			// The old 'last' becomes the new 'second last'
			lastSeenIndex := lastSeen[value]
			secondLastSeen[value] = lastSeenIndex
			// The current index 'i' becomes the new 'last'
			lastSeen[value] = i

		} else if _, ok := lastSeen[value]; ok {
			// Case 2: Second occurrence found.
			// Promote the first index (from lastSeen) to 'secondLastSeen'
			secondLastSeen[value] = lastSeen[value]
			// Update 'lastSeen' to the current index
			lastSeen[value] = i

		} else {
			// Case 3: First occurrence found.
			// Just record its index in 'lastSeen'
			lastSeen[value] = i
		}
	}

	// If minDistance is still MaxInt32, no triplet was found, return -1.
	if minDistance == math.MaxInt32 {
		return -1
	}
	return minDistance
}