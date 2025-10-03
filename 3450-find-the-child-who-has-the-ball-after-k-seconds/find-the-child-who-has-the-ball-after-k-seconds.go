func numberOfChild(n int, k int) int {
    // Calculate the effective position within the cycle
    pos := k % (2*n - 2)
    
    // Determine the direction of movement
    if pos < n {
        return pos
    }
    return 2*n - 2 - pos
}
