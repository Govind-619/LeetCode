func findRotation(mat [][]int, target [][]int) bool {
    n := len(mat)

    for r := 0; r < 4; r++ {
        match := true

        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                if mat[i][j] != target[i][j] {
                    match = false
                    break
                }
            }
            if !match {
                break
            }
        }

        if match {
            return true
        }

        rotate(mat) // rotate for next iteration
    }

    return false
}

func rotate(matrix [][]int) {
    n := len(matrix)

    // Step 1: Transpose
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }

    // Step 2: Reverse each row
    for i := 0; i < n; i++ {
        for j := 0; j < n/2; j++ {
            matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
        }
    }
}