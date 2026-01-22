func minimumPairRemoval(nums []int) int {
    if slices.IsSorted(nums){
        return 0
    }

    idx, sum := 0, nums[0] + nums[1]
    for i := 1; i < len(nums) - 1; i++ {
        tempSum := nums[i] + nums[i+1]
        if tempSum < sum {
            idx, sum = i, tempSum
        }
    }

    nums[idx] = sum
    return minimumPairRemoval(append(nums[:idx+1], nums[idx+2:]...)) + 1
}