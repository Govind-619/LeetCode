func concatenatedBinary(n int) int {
    const MOD int = 1_000_000_007
    result := 0
    bitLength := 0

    for i := 1; i <= n; i++ {
        if i&(i-1) == 0 {
            bitLength++
        }
        result = ((result << bitLength) + i) % MOD
    }

    return result
}