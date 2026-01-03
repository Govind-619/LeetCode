func numOfWays(n int) int {
    const MOD = 1000000007
    A, B := 6, 6

    for i := 2; i <= n; i++ {
        newA := (3*A + 2*B) % MOD
        newB := (2*A + 2*B) % MOD
        A, B = newA, newB
    }

    return (A + B) % MOD
}
