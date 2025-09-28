func largestPerimeter(nums []int) int {
    sort.Ints(nums)
    n:= len(nums)-1
    for i:= n; i>1; i--{
        if nums[i-1] + nums[i-2] > nums[i]{
            return nums[i-1] + nums[i-2] + nums[i]
        }
    }
    return 0
}