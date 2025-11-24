func prefixesDivBy5(nums []int) []bool {
    ans := make([]bool, len(nums))
    var num int 
    for i := 0; i < len(nums); i++ {
        num = (num << 1 | nums[i]) % 5
        ans[i] = num == 0
    }
    return ans
}