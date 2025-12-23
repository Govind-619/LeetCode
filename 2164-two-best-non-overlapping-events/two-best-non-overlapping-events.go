import (
	"sort"
)

func maxTwoEvents(events [][]int) int {
	// Step 1: Sort events by their end times
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	n := len(events)
	maxValues := make([]int, n)
	maxValues[0] = events[0][2] // Value of the first event

	// Fill the running max array
	for i := 1; i < n; i++ {
		maxValues[i] = max(maxValues[i-1], events[i][2])
	}

	maxSum := 0

	// Step 3: Iterate through the events and use binary search
	for i := 0; i < n; i++ {
		start := events[i][0]
		value := events[i][2]

		// Find the latest event that ends before the current event starts
		idx := binarySearch(events, start-1)

		currentSum := value
		if idx != -1 {
			currentSum += maxValues[idx]
		}

		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

// Helper: Binary search to find the latest event that ends <= targetEnd
func binarySearch(events [][]int, targetEnd int) int {
	start, end, res := 0, len(events)-1, -1

	for start <= end {
		mid := start + (end-start)/2
		if events[mid][1] <= targetEnd {
			res = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return res
}

// Utility: Max function
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}