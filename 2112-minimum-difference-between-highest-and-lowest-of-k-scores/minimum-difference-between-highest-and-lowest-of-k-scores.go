func minimumDifference(nums []int, k int) int {
    if len(nums)==1{
        return 0
    }
    sort.Ints(nums)
    mind:= math.MaxInt
    for i:=0; i<=len(nums)-k; i++{
       mind =min(mind,nums[i+k-1]-nums[i])
    }
    return mind
}