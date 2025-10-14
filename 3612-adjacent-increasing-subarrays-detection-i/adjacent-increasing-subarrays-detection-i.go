func hasIncreasingSubarrays(nums []int, k int) bool {
    n := len(nums)
    if n < 2*k {
        return false
    }

    inc := make([]int, n)
    inc[0] = 1

    // Step 1: Calculate lengths of current increasing streaks ending at i
    for i := 1; i < n; i++ {
        if nums[i] > nums[i-1] {
            inc[i] = inc[i-1] + 1
        } else {
            inc[i] = 1
        }
    }

    // Step 2: Check for two adjacent runs of length k
    for i := k; i < n; i++ {
        // inc[i] >= k → last k elements form an increasing subarray
        // check if the previous k elements also form an increasing subarray
        if inc[i] >= k && inc[i-k] >= k {
            return true
        }
    }
    return false
}
