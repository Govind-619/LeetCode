func climbStairs(n int, costs []int) int {
    costs = append([]int{0}, costs...)
    a, b, c := costs[n], math.MaxInt/2, math.MaxInt/2  // dp[i + 1], dp[i + 2], dp[i + 3] for i = n - 1
    for i := n - 1; i >= 0; i-- {
        a, b, c = costs[i]+min(a+1, b+4, c+9), a, b
    }
    return a
}