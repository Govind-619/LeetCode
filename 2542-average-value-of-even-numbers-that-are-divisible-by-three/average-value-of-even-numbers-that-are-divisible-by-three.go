func averageValue(nums []int) int {
    sum,count:=0,0
    for _,v:= range nums{
        if v%6==0{
            count++
            sum+=v
        }
    }
    if count!=0{
        return int(sum/count)
    }
     return 0
}