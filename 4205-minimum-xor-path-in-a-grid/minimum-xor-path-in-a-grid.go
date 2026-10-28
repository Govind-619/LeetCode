func minCost(grid [][]int) int {
    n := len(grid[0])
    dp := make([][1024]bool, n)  // dp[j][x]: XOR x can reach column j of the row above
    dp[0][0] = true
    for _, row := range grid {
        for j, v := range row {
            var next [1024]bool
            for x := 0; x < 1024; x++ {
                if dp[j][x] || j > 0 && dp[j-1][x] {  // from above or from the left
                    next[x^v] = true
                }
            }
            dp[j] = next
        }
    }
    x := 0
    for !dp[n-1][x] {
        x++
    }
    return x
}