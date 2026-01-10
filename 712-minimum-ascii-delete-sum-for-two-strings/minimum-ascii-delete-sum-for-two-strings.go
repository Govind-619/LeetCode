func minimumDeleteSum(s1 string, s2 string) int {
    
    n, m := len(s1), len(s2)
    dp := make([][]int, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([]int, m+1)
    }

    for i := 1; i <= n; i++ {
        for j := 1; j <= m; j++ {
            if s1[i-1] == s2[j-1] {
                dp[i][j] = dp[i-1][j-1] + int(s1[i-1])
            } else {
                if dp[i-1][j] > dp[i][j-1] {
                    dp[i][j] = dp[i-1][j]
                } else {
                    dp[i][j] = dp[i][j-1]
                }
            }
        }
    }

    kept := dp[n][m]
    result := asciiSum(s1) + asciiSum(s2) - 2*kept

    return result
}

func asciiSum(s string) int {
    sum := 0
    for i := 0; i < len(s); i++ {
        sum += int(s[i])
    }
    return sum
}
