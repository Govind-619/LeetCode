func closestTarget(words []string, target string, startIndex int) int {
    n := len(words)
    // Initialize with a value larger than any possible distance
    minDist := n

    found := false
    for i, v := range words {
        if v == target {
            found = true
            // Standard absolute distance
            dist := i - startIndex
            if dist < 0 {
                dist = -dist
            }
            
            // Circular distance: either direct or wrapping around the edge
            // n - dist represents the distance going the "other way" around the circle
            currentDist := dist
            if n - dist < currentDist {
                currentDist = n - dist
            }
            
            if currentDist < minDist {
                minDist = currentDist
            }
        }
    }

    if !found {
        return -1
    }
    return minDist
}