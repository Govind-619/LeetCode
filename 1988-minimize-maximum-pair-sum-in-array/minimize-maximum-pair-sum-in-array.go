func minPairSum(nums []int) int {
    sort.Ints(nums)
    sum:= []int{}
    for i, j:= 0, len(nums)-1; i<j; i,j= i+1, j-1{
        sum= append(sum,nums[i]+nums[j])
    }
    sort.Ints(sum)
    return sum[len(sum)-1]
}