func hasSameDigits(s string) bool {
    r:= []rune(s)
    for len(r)>2{
        for i:= 0; i<len(r)-1; i++{
            r[i]= (r[i]+ r[i+1])%10
        }
        r= r[:len(r)-1]
    }
    if r[0]==r[1]{
        return true
    }
    return false
}
