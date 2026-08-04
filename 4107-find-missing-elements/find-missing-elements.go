func findMissingElements(nums []int) []int {
    sort.Ints(nums)
    res := []int{}
    z := 0
    start, stop := nums[0], nums[len(nums)-1]
    
    for i := start; i < stop; i++ {
        if i == nums[z] {
            // The number is present in the array, move to the next element in nums
            z++
        } else {
            // The number is missing, add it to the result
            // Do not increment z, because we still need to compare nums[z] against the next i
            res = append(res, i)
        }
    }
    
    return res
}