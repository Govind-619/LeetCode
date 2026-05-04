func rotate(matrix [][]int) {
    n := len(matrix)
    
    // 1. Transpose: Swap elements across the main diagonal
    // Only iterate where j > i to avoid swapping back
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    
    // 2. Reverse each row: This completes the 90-degree clockwise rotation
    for i := 0; i < n; i++ {
        for k, j := 0, n-1; k < j; k, j = k+1, j-1 {
            matrix[i][k], matrix[i][j] = matrix[i][j], matrix[i][k]
        }
    }
}