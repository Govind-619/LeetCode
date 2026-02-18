func hasAlternatingBits(n int) bool {
    c:= fmt.Sprintf("%b",n)
    fmt.Println(c)
    z:=rune(c[0])
    for _,v:= range c[1:]{
        if v==z{
            return false
        }
        z=v
    }
    return true
}