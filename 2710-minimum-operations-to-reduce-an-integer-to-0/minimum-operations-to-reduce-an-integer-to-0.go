func Log2(n int) int {
    power := 0
    for n > 1 {
        n >>= 1
        power++
    }
    return power
}

func minOperations(n int) int {
    cnt := 0
    for n > 0 {
        cnt++
        pow2 := Log2(n)
        n = min(n - (1 << pow2), (1 << (pow2 + 1)) - n)
    }
    return cnt
}