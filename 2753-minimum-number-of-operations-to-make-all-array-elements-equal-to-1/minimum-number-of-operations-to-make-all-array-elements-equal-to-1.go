func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func minOperations(nums []int) int {
    n := len(nums)
    g := nums[0]
    for _, v := range nums {
        g = gcd(g, v)
    }
    if g != 1 {
        return -1
    }

    // If we already have 1s
    ones := 0
    for _, v := range nums {
        if v == 1 {
            ones++
        }
    }
    if ones > 0 {
        return n - ones
    }

    // Find smallest subarray with gcd == 1
    minLen := n
    for i := 0; i < n; i++ {
        g := nums[i]
        for j := i + 1; j < n; j++ {
            g = gcd(g, nums[j])
            if g == 1 {
                if j-i < minLen {
                    minLen = j - i
                }
                break
            }
        }
    }

    return minLen + n - 1
}
