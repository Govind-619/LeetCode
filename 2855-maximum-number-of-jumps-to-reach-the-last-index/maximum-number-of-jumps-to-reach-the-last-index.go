func maximumJumps(nums []int, target int) int {
    
    n := len(nums)

    dp := make([]int, n)

    for i := 0; i < n; i++ {
        dp[i] = -1
    }

    dp[0] = 0

    for i := 0; i < n; i++ {

        if dp[i] == -1 {
            continue
        }

        for j := i + 1; j < n; j++ {

            diff := nums[j] - nums[i]

            if diff >= -target && diff <= target {

                if dp[i]+1 > dp[j] {
                    dp[j] = dp[i] + 1
                }
            }
        }
    }

    return dp[n-1]
}