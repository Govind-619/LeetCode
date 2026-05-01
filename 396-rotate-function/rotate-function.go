func maxRotateFunction(nums []int) int {
    n := len(nums)
    sum := 0
    f := 0
    
    for i := 0; i < n; i++ {
        sum += nums[i]
        f += i * nums[i]
    }
    
    maxVal := f
    
    // Calculate subsequent F(k) using the O(1) formula
    for i := n - 1; i > 0; i-- {
        f = f + sum - n*nums[i]
        if f > maxVal {
            maxVal = f
        }
    }
    
    return maxVal
}