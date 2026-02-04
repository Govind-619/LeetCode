func alternatingSum(nums []int) int {
    odd,even:=0,0
    for i:= range nums{
        if i%2==0{
            odd+=nums[i]
        }else{
            even+=nums[i]
        }
    }
    return odd-even
}