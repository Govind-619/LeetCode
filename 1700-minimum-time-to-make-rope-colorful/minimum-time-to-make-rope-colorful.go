func minCost(colors string, neededTime []int) int {
    total := 0
    maxTime := neededTime[0]
    
    for i := 1; i < len(colors); i++ {
        if colors[i] == colors[i-1] {
            total += min(maxTime, neededTime[i])
            // keep the larger time for the next comparison
            if neededTime[i] > maxTime {
                maxTime = neededTime[i]
            }
        } else {
            maxTime = neededTime[i]
        }
    }
    return total
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
