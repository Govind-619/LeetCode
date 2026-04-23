func distance(nums []int) []int64 {
    n := len(nums)
    res := make([]int64, n)
    
    // Group indices by their value in nums
    groups := make(map[int][]int)
    for i, v := range nums {
        groups[v] = append(groups[v], i)
    }
    
    // Process each group of indices
    for _, indices := range groups {
        m := len(indices)
        if m <= 1 {
            continue
        }
        
        // Calculate total sum of indices for this group to help find suffix sums
        var totalSum int64
        for _, idx := range indices {
            totalSum += int64(idx)
        }
        
        var prefixSum int64
        for i, idx := range indices {
            idx64 := int64(idx)
            i64 := int64(i)
            m64 := int64(m)
            
            // Sum of distances to the left: (count * currentIdx) - prefixSum
            leftSide := i64*idx64 - prefixSum
            
            // Sum of distances to the right: suffixSum - (count * currentIdx)
            // suffixSum = totalSum - prefixSum - currentIdx
            suffixSum := totalSum - prefixSum - idx64
            rightSide := suffixSum - (m64-1-i64)*idx64
            
            res[idx] = leftSide + rightSide
            
            // Update prefixSum for the next index in the group
            prefixSum += idx64
        }
    }
    
    return res
}