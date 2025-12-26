func maximumHappinessSum(happiness []int, k int) int64 {
    sort.Ints(happiness)
    n:= len(happiness)
    res:= 0
    count:= 0
    for i:= n-1;i>=n-k; i--{
        if happiness[i]-count>0{
            res+= happiness[i]-count
        }
        count++
    } 
    return int64(res)
}