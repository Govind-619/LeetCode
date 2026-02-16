func reverseBits(n int) int {
    ans, mult := 0, 1

    for i := 31; i >= 0; i-- {
        bit := (n >> i) & 1
        ans += bit * mult
        mult <<= 1
    }

    return ans
}