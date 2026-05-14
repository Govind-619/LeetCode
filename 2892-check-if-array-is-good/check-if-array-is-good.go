func isGood(nums []int) bool {

    sort.Ints(nums)

    n := len(nums)

    mx := nums[n-1]

    if n != mx+1 {
        return false
    }

    for i := 0; i < n-1; i++ {

        if nums[i] != i+1 {
            return false
        }
    }

    return nums[n-1] == mx
}