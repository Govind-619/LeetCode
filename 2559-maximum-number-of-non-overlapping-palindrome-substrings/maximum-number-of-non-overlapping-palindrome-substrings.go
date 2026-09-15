func maxPalindromes(s string, k int) int {
    n := len(s)
    isPal := make([][]bool, n)
    for i := range isPal {
        isPal[i] = make([]bool, n)
    }
    
    for i := 0; i < n; i++ {
        isPal[i][i] = true
    }
    
    for i := 0; i+1 < n; i++ {
        if s[i] == s[i+1] {
            isPal[i][i+1] = true
        }
    }
    
    for length := 3; length <= n; length++ {
        for i := 0; i+length-1 < n; i++ {
            j := i + length - 1
            if s[i] == s[j] && isPal[i+1][j-1] {
                isPal[i][j] = true
            }
        }
    }
    
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = dp[i-1]
        for j := 0; j <= i-k; j++ {
            if isPal[j][i-1] {
                if dp[j]+1 > dp[i] {
                    dp[i] = dp[j] + 1
                }
            }
        }
    }
    return dp[n]
}