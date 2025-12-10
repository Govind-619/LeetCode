func countPermutations(comp []int) int {
    n := len(comp)
    for i := 1; i < n; i++ {
        if comp[i] <= comp[0] {
            return 0
        }
    }
    
    mod := int(1e9+7)
    prod := 1
    for count := n-1; count > 1; count-- {
        prod = (prod * count) % mod
    }

    return prod
}