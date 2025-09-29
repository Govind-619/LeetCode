import "fmt"

func minScoreTriangulation(values []int) int {
    n := len(values)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }

    for length := 2; length < n; length++ {
        for i := 0; i+length < n; i++ {
            j := i + length
            dp[i][j] = 1<<31 - 1
            for k := i + 1; k < j; k++ {
                cost := dp[i][k] + dp[k][j] + values[i]*values[j]*values[k]
                if cost < dp[i][j] {
                    dp[i][j] = cost
                }
            }
        }
    }
    return dp[0][n-1]
}