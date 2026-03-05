func minOperations(s string) int {
    res := 0
    for i, l := range s {
        res += (i & 1) ^ int(l-'0')
    }
    return min(res, len(s)-res)
}