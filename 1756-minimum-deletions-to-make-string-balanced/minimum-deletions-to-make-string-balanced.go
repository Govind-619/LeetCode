func minimumDeletions(s string) int {
    r,b:=0,0

    for _,v:= range s{
        if v=='b'{
            b++
        }else{
            r= min(r+1,b)
        }
    } 

    return r
}