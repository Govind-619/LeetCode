func maxOperations(s string) int {
    ones := 0
    ans := 0
    n := len(s)

    for i := 0; i < n; i++ {
        if s[i] == '1' {
            ones++
        } else if i == n-1 || s[i+1] == '1' {
            ans += ones
        }
    }

    return ans
}
