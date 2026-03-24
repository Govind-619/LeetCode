func constructProductMatrix(grid [][]int) [][]int {
    row, col := len(grid), len(grid[0])
    mod := 12345

    // Flatten
    n := row * col
    arr := make([]int, 0, n)
    for _, r := range grid {
        arr = append(arr, r...)
    }

    // Prefix products
    prefix := make([]int, n)
    prefix[0] = 1
    for i := 1; i < n; i++ {
        prefix[i] = (prefix[i-1] * arr[i-1]) % mod
    }

    // Suffix products
    suffix := make([]int, n)
    suffix[n-1] = 1
    for i := n - 2; i >= 0; i-- {
        suffix[i] = (suffix[i+1] * arr[i+1]) % mod
    }

    // Result array
    resArr := make([]int, n)
    for i := 0; i < n; i++ {
        resArr[i] = (prefix[i] * suffix[i]) % mod
    }

    // Convert back to 2D
    res := make([][]int, row)
    idx := 0
    for i := 0; i < row; i++ {
        res[i] = make([]int, col)
        for j := 0; j < col; j++ {
            res[i][j] = resArr[idx]
            idx++
        }
    }

    return res
}