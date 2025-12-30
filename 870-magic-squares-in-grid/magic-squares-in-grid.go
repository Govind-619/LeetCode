func numMagicSquaresInside(grid [][]int) int {
    result := 0
    row, col := len(grid), len(grid[0])
    for i := 0 ; i < row-2; i++{
        for j := 0; j < col-2; j++{
            // create 3 * 3 matrix from index i,j
            matrix := createMatrix(grid,i,j)
            if isMagicSquare(matrix) {
                result++
            }
        }
    }

    return result
}

func createMatrix(grid [][]int, i, j int)[3][3]int{
    matrix := [3][3]int{}
    for x := 0 ; x <= 2 ; x++{
        matrix[x] = [3]int{}
        for y := 0 ; y <= 2 ; y++{
            matrix[x][y] = grid[x+i][y+j]
        }
    }
    return matrix
}

func isMagicSquare(matrix [3][3]int)bool{
    counter := [10]int{}
    for _, row:= range matrix{
        for _, v := range row{
            if v > 9  || v <= 0{
                return false
            }
            counter[v]++
            if counter[v] > 1 {
                return false
            }
        }
    }
    magicNumber := matrix[0][0]+matrix[0][1]+matrix[0][2]
    magicSquares := make([]int,0)
    magicSquares = append(magicSquares, matrix[1][0]+matrix[1][1]+matrix[1][2])
    magicSquares = append(magicSquares, matrix[2][0]+matrix[2][1]+matrix[2][2])
    magicSquares = append(magicSquares, matrix[0][0]+matrix[1][0]+matrix[2][0])
    magicSquares = append(magicSquares, matrix[0][1]+matrix[1][1]+matrix[2][1])
    magicSquares = append(magicSquares, matrix[0][2]+matrix[1][2]+matrix[2][2])
    magicSquares = append(magicSquares, matrix[0][0]+matrix[1][1]+matrix[2][2])
    magicSquares = append(magicSquares, matrix[0][2]+matrix[1][1]+matrix[2][0])

    for _, v := range magicSquares{
        if magicNumber != v {
            return false
        }
    }

    return true
}