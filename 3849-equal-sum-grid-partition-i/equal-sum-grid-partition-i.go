func canPartitionGrid(grid [][]int) bool {
    m, n := len(grid), len(grid[0])
    
    total := 0
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            total += grid[i][j]
        }
    }
    
    if total%2 != 0 {
        return false
    }
    
    target := total / 2
    
    rowSum := 0
    for i := 0; i < m-1; i++ {
        for j := 0; j < n; j++ {
            rowSum += grid[i][j]
        }
        if rowSum == target {
            return true
        }
    }
    
    colSum := make([]int, n)
    for j := 0; j < n; j++ {
        for i := 0; i < m; i++ {
            colSum[j] += grid[i][j]
        }
    }
    
    curr := 0
    for j := 0; j < n-1; j++ {
        curr += colSum[j]
        if curr == target {
            return true
        }
    }
    
    return false
}