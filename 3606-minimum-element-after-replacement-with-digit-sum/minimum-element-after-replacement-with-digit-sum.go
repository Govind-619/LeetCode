func minElement(nums []int) int {
    a:= []int{}
    for _,v:= range nums{
        if v/10==0{
            a= append(a,v)
        }else{
            s:=0
            for v>0{
                s+=v%10
                v/=10
            }
            a= append(a,s)
        }
    }
    sort.Ints(a)
    return a[0]
}