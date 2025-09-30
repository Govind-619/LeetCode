func triangleType(nums []int) string {
    sort.Ints(nums)
    switch{
        case nums[0]==nums[1] && nums[1]==nums[2]:
        return "equilateral"
        case nums[0]==nums[1] && nums[0]+nums[1]>nums[2]|| nums[1]==nums[2] ||nums[0]==nums[2]:
        return "isosceles"
        case nums[0]+nums[1]>nums[2]:
        return "scalene"
        default:
        return "none"
    }
}