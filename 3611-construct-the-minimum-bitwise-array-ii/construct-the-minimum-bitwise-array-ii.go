func minBitwiseArray(nums []int) []int {
    res := make([]int, len(nums))
    for i, n := range nums {
        res[i] = findNumOptimized(n)
    }
    return res
}

func findNumOptimized(x int) int {
    // Only even prime is 2 → impossible
    if x%2 == 0 {
        return -1
    }
    b := 1
    for x & b == b && x > b { 
        b = b<<1 
    }
    if x <= b { 
        return x>>1 
    }

    return x-(b>>1)
}
