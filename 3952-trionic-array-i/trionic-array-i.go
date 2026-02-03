func isTrionic(nums []int) bool {
    n := len(nums)
    i := 1
    for i < n && nums[i] > nums[i-1] {
        i++
    }
    if i == 1 || i == n {
        return false
    }
    for i < n && nums[i] < nums[i-1] {
        i++
    }
    if i == 2 || i == n {
        return false
    }
    for i < n && nums[i] > nums[i-1] {
        i++
    }
    return i == n
}