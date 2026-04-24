func furthestDistanceFromOrigin(moves string) int {
    l,r,n:=0,0,0
    for _,v:= range moves{
        if v=='L'{
            l++
        }else if v=='R'{
            r++
        }else{
            n++
        }
    }
    if l>r{
        return (l+n)-r
    }else if r>l{
        return (r+n)-l
    }else{
        return n
    }
}