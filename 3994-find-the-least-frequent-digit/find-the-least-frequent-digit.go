type val struct{
    v int
    f int
}
func getLeastFrequentDigit(n int) int {
    m:= make(map[int]int)
    for n>0{
        m[n%10]++
        n/=10
    }
    res:= []val{}
    for k,v:= range m{
        res= append(res,val{k,v})
    }
    sort.Slice(res,func(i,j int)bool{
       if res[i].f!=res[j].f{
        return res[i].f<res[j].f
       }else{
        return res[i].v<res[j].v
       }
    })
    return res[0].v
}