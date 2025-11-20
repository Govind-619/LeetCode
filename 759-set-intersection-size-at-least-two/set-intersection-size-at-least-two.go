func intersectionSizeTwo(intervals [][]int) int {
    // Sort by end point, then by start point (descending)
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][1] != intervals[j][1] {
            return intervals[i][1] < intervals[j][1]
        }
        return intervals[i][0] > intervals[j][0]
    })
    
    result := 0
    p1, p2 := -1, -1  // Track the two largest elements
    
    for _, interval := range intervals {
        start, end := interval[0], interval[1]
        
        // Case 1: Both p1 and p2 are in the interval
        if start <= p1 {
            continue
        }
        // Case 2: Only p2 is in the interval
        if start <= p2 {
            p1 = p2
            p2 = end
            result++
        } else {
            // Case 3: Neither is in the interval
            p1 = end - 1
            p2 = end
            result += 2
        }
    }
    
    return result
}