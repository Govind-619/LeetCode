func countPartitions(nums []int, k int) int {
    const MOD = 1_000_000_007
    n := len(nums)

    dp := make([]int, n+1)
    dp[0] = 1

    mx := make([]int, n)
    mn := make([]int, n)
    mxl, mxr := 0, 0
    mnl, mnr := 0, 0

    l := 0
    sum := 0

    for r := 0; r < n; r++ {
        for mxl < mxr && nums[mx[mxr-1]] <= nums[r] {
            mxr--
        }
        for mnl < mnr && nums[mn[mnr-1]] >= nums[r] {
            mnr--
        }
        mx[mxr] = r
        mxr++
        mn[mnr] = r
        mnr++

        for nums[mx[mxl]]-nums[mn[mnl]] > k {
            if mx[mxl] == l {
                mxl++
            }
            if mn[mnl] == l {
                mnl++
            }
            sum = (sum - dp[l] + MOD) % MOD
            l++
        }

        sum = (sum + dp[r]) % MOD
        dp[r+1] = sum
    }

    return dp[n]
}