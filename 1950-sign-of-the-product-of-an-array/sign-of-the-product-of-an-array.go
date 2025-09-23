func arraySign(nums []int) int {
    neg:= 0
    for _, v:= range nums{
        if v<0{
            neg++
        }else if v==0{
            return 0
        } 
    }
    if neg%2==0{
        return 1
    }
    return -1
}