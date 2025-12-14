func numberOfWays(corridor string) int {
    chairInds := []int{}
    for i, ch := range corridor {
        if ch == 'S' {
            chairInds = append(chairInds, i)
        } 
    }

    if len(chairInds) % 2 != 0 || len(chairInds) == 0 {
        return 0
    }

    out := 1
    mod := int(1e9+7)
    for i := 1; i < len(chairInds)-1; i += 2 {
        out = (out * (chairInds[i+1] - chairInds[i])) % mod
    }

    return out
}