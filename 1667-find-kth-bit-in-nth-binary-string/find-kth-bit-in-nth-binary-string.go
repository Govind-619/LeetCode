func findKthBit(n int, k int) byte {
    inverted := false
    for n > 1 {
        mid := 1 << uint(n-1)
        if k == mid {
            if inverted {
                return '0'
            }
            return '1'
        }
        if k > mid {
            k = 2*mid-k
            inverted = !inverted
        }
        n--
    }
    if inverted {
        return '1'
    }
    return '0'
}