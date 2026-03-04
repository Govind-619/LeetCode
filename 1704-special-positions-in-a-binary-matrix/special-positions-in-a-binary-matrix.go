
func numSpecial(mat [][]int) int {
    m,n := len(mat), len(mat[0])
    for i := 0; i < m; i++ {
        // check row validity
        count := 0
        for j := 0; j < n && count < 2; j++ {
            if mat[i][j] == 1 {
                count++
            }
        }
        // if invalid mark them with 2
        if count >= 2 {
            for j := 0; j < n; j++ {
                if mat[i][j] == 1 {
                    mat[i][j] = 2
                }
            }            
        }
    }

    // check cols 
    res := 0
    for j := 0; j < n; j++ {
        count := 0
        x,y := -1,-1
        for i := 0; i < m && count < 2; i++ {
            if mat[i][j] >= 1 {
                x,y = i,j
                count++
            }
        }
        if count == 1 && mat[x][y] == 1 {
            res++
        }
    }

    return res
}