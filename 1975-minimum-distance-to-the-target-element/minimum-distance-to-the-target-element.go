func getMinDistance(nums []int, target int, start int) int {
    vals:= math.MaxInt
    for i,v:= range nums{
        if v== target{
            if abs(i-start)<vals{
                vals=abs(i-start)
            }
        }
    }
    return vals
}

func abs(x int)int{
    if x<0{
        return -x
    }
    return x
}