func minMoves(nums []int) int {
    if len(nums)==1{
        return 0
    }
    max:= nums[0]
    count:=0
    for _,v:= range nums[1:]{
        if v>max{
            max=v
        }
    }
    for _,v:= range nums{
        if v<max{
            count+= max-v
        }
    }
    return count   
}