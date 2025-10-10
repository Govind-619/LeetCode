func maximumEnergy(energy []int, k int) int {
    n := len(energy)
    dp := make([]int, n)

    // Start from the last index and move backwards
    for i := n - 1; i >= 0; i-- {
        dp[i] = energy[i]
        if i+k < n {
            dp[i] += dp[i+k] // jump forward by k if possible
        }
    }

    // Find the maximum among all dp values
    res := dp[0]
    for i := 1; i < n; i++ {
        if dp[i] > res {
            res = dp[i]
        }
    }

    return res
}
