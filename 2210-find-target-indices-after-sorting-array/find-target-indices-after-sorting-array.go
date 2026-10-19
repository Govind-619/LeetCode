func targetIndices(nums []int, target int) []int {
    sort.Ints(nums)
    res:=[]int{}
    for i,v:= range nums{
        if v==target{
            res= append(res,i)
        }  
    }
    return res
}