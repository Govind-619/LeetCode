func getSneakyNumbers(nums []int) []int {
    out:= []int{}
    m:=make(map[int]bool)
    for _,v := range nums{
        if !m[v]{
            m[v]=true
        }else{
            out =append(out,v)
        }
    }
return out
    
}