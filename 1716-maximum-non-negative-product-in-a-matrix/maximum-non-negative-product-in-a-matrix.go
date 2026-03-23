// keep the maximum positive and negative product for each cell
type top struct { pos,neg int64 }
const (
    z64 = int64(0)
    o64 = int64(1)
    mod = 1000000007
)
func maxProductPath(grid [][]int) int {
    // one row of greatest products
    tops := make([]top, len(grid[0]))
    tops[0] = top{ o64, z64 }
    // check if there is a zero in the matrix
    z := false
    for i:=0; i<len(grid); i++ {
        if grid[i][0] == 0 { z = true }
        // element in the first column can only be reached from above
        tops[0] = update(grid[i][0], tops[0], top{ z64, z64 })
        for j:=1; j<len(grid[0]); j++ {
            if grid[i][j] == 0 { z = true }
            // other elements can be reached from above or from the left
            tops[j] = update(grid[i][j], tops[j], tops[j-1])
        }
    }
    r := tops[len(tops)-1].pos
    // no positive product and no zero in the matrix: no valid solution
    if r == z64 && !z { return -1 }
    return int(r % mod)
}
// calculate new maximum products with adjacent cells multiplied by current value
func update(v int, t1, t2 top) top {
    // take the greater negative and positive products along the path
    ps, ng := max(t1.pos, t2.pos), max(t1.neg, t2.neg)
    if ps > -1 { 
        // multiply positive product by current cell
        if v < 0 { ps *= int64(-v) } else { ps *= int64(v) }
    }
    if ng > -1 { 
        // multiply negative product by current cell
        if v < 0 { ng *= int64(-v) } else { ng *= int64(v) }
    }
    // if the current value is negative, swap the positive and negative values
    if v < 0 { ps, ng = ng, ps }
    return top{ ps, ng }
}