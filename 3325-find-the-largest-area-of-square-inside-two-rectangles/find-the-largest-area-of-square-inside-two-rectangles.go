func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
    n := len(bottomLeft)    
    ans := int64(0)
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            blX := max(bottomLeft[i][0], bottomLeft[j][0])
            blY := max(bottomLeft[i][1], bottomLeft[j][1])
            trX := min(topRight[i][0], topRight[j][0])
            trY := min(topRight[i][1], topRight[j][1])
            
            if blX < trX && blY < trY {
                side := int64(min(trX - blX, trY - blY))
                ans = max(ans, side * side)
            }
        }
    }

    return ans
}