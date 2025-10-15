func maxIncreasingSubarrays(nums []int) int {
    ans, pre, cur := 0, 0, 0
    n := len(nums)
    for i, x := range nums {
        cur++
        // if we are at the end or next is not greater
        if i == n-1 || x >= nums[i+1] {
            // update answer
            if cur/2 > ans {
                ans = cur/2
            }
            m := pre
            if cur < m {
                m = cur
            }
            if m > ans {
                ans = m
            }
            pre = cur
            cur = 0
        }
    }
    return ans
}
