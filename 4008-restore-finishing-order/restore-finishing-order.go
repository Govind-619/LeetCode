func recoverOrder(order []int, friends []int) []int {
    m:= make(map[int]bool)
    for _,v:= range friends{
        m[v]= true
    }
    res:= []int{}
     for _,v:= range order{
        if m[v]{
            res= append(res,v)
        }
     }
    return res
}