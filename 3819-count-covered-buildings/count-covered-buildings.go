func countCoveredBuildings(n int, buildings [][]int) int {
    maxX := make([]int, n+1)
    minX := make([]int, n+1)
    maxY := make([]int, n+1)
    minY := make([]int, n+1)

    for i := 0; i < n+1; i++ {
        minX[i] = math.MaxInt
        minY[i] = math.MaxInt
    }

    for _, b := range buildings {
        x, y := b[0], b[1]
        maxX[x] = max(maxX[x], y)
        minX[x] = min(minX[x], y)
        maxY[y] = max(maxY[y], x)
        minY[y] = min(minY[y], x)
    }

    ans := 0
    for _, b := range buildings {
        x, y := b[0], b[1]
        if y > minX[x] && y < maxX[x] && x > minY[y] && x < maxY[y] {
            ans++
        }
    }

    return ans
}