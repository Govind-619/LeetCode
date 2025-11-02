func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
    grid := make([][]int, m)
    for i := range grid {
        grid[i] = make([]int, n)
    }

    // Mark guards as 1 and walls as 2
    for _, g := range guards {
        grid[g[0]][g[1]] = 1
    }
    for _, w := range walls {
        grid[w[0]][w[1]] = 2
    }

    dirs := [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}} // down, up, right, left

    for _, g := range guards {
        for _, d := range dirs {
            x, y := g[0]+d[0], g[1]+d[1]
            for x >= 0 && x < m && y >= 0 && y < n && grid[x][y] != 1 && grid[x][y] != 2 {
                // Mark as guarded (3)
                if grid[x][y] == 0 {
                    grid[x][y] = 3
                }
                x += d[0]
                y += d[1]
            }
        }
    }

    // Count unguarded cells
    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 {
                count++
            }
        }
    }

    return count
}
