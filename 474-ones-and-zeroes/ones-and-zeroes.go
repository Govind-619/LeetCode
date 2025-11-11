func findMaxForm(strs []string, m int, n int) int {
    list := make([][2]int, len(strs))
    for i := 0; i < len(strs); i++ {
        l := len(strs[i])
        list[i][0] = strings.Count(strs[i], "0")
        list[i][1] = l - list[i][0]
    }
    dp := make([][]int, m + 1) 
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n + 1)
    }
    dp[0][0] = 0
    
    for k := 0; k < len(list); k++ {
        for i := m; i >= 0; i-- {
            for j := n; j >= 0; j-- {
                if i >= list[k][0] && j >= list[k][1] {
                    dp[i][j] = max(dp[i][j], dp[i - list[k][0]][j - list[k][1]] + 1)
                } else {
                    break
                }
            }
        }
    }
    return dp[m][n]
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}