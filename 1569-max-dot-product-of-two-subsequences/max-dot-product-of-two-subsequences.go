func maxDotProduct(nums1 []int, nums2 []int) int {
    n, m := len(nums1), len(nums2)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, m)
        for j := range memo[i] {
            memo[i][j] = math.MinInt32
        }
    }

    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if i == n || j == m {
            return math.MinInt32 // we use max(0, dfs(i+1, j+1)) later in take, thus if we never took it will be math.Int32 so we will throw it away
        }
        if memo[i][j] != math.MinInt32 {
            return memo[i][j]
        }

        take := nums1[i] * nums2[j] + max(0, dfs(i+1, j+1)) // we can throw away negative products
        skip1 := dfs(i+1, j)
        skip2 := dfs(i, j+1)

        memo[i][j] = max(take, skip1,skip2)
        return memo[i][j]
    }

    return dfs(0, 0)
}