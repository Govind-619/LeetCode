func largestMagicSquare(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    // start with the largest possible size first
    for z:=min(n,m); z>1; z-- {
        // check at every possible top-left corner
        for x:=0; x<=n-z; x++ {
            for y:=0; y<=m-z; y++ {
                // if magic square is found, it is the largest
                if magic(grid, x, y, z) { return z }
            }
        }
    }
    // every single cell is a 1x1 magic square
    return 1
}
func magic(grid [][]int, x int, y int, z int) bool {
    // compare diagonals
    t1, t2 := 0, 0
    for i:=range z { 
        t1 += grid[x+i][y+i] 
        t2 += grid[x+z-1-i][y+i]
    }
    // check each row
    if t1 != t2 { return false }
    for i:=range z {
        t2 = 0
        for j:=range z { t2 += grid[x+i][y+j] }
        if t1 != t2 { return false }
    }
    // check each column
    for j:=range z {
        t2 = 0
        for i:=range z { t2 += grid[x+i][y+j] }
        if t1 != t2 { return false }
    }
    return true
}