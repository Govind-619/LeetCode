func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
    for i := 0; i < k/2; i++ {               // swap rows inside submatrix
        for j := 0; j < k; j++ {             // iterate columns
            grid[x+i][y+j], grid[x+k-1-i][y+j] = grid[x+k-1-i][y+j], grid[x+i][y+j]
        }
    }
    return grid
}