func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
    sort.Ints(hBars)
    sort.Ints(vBars)
    h := min(biggestHole(hBars, n), biggestHole(vBars, m))
    return h * h
}

func biggestHole(bars []int, size int) int {
    h, mx := 1, 1
    for i, b := range bars {
        if i == 0 || bars[i] == bars[i-1]+1 {
            if b > 1 && b < size+2 { 
                h++ 
            }
        } else {
            if h > mx {
                mx = h 
            }
            h = 1
            if b > 1 && b < size+2 {
                h++ 
            }
        }
    }
    if h > mx { 
        mx = h 
    }
    return mx
}