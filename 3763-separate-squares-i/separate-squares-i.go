package main

func helper(line float64, squares [][]int) float64 {
    var aAbove, aBelow float64
    n := len(squares)
    
    for i := 0; i < n; i++ {
        y, l := squares[i][1], squares[i][2]
        total := float64(l * l)
        
        if line <= float64(y) {
            aAbove += total
        } else if line >= float64(y+l) {
            aBelow += total
        } else {
            aboveHeight := float64(y+l) - line
            belowHeight := line - float64(y)
            aAbove += float64(l) * aboveHeight
            aBelow += float64(l) * belowHeight
        }
    }
    
    return aAbove - aBelow
}

func separateSquares(squares [][]int) float64 {
    lo, hi := 0.0, 2e9
    
    for i := 0; i < 60; i++ {
        mid := (lo + hi) / 2.0
        diff := helper(mid, squares)
        
        if diff > 0 {
            lo = mid
        } else {
            hi = mid
        }
    }
    
    return hi
}