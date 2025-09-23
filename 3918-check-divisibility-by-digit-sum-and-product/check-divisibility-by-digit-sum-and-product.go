func checkDivisibility(n int) bool {
    s,p:= sp(n)

    if n% (s+p)==0{
        return true
    }
    return false
}

func sp(x int)(int, int){
    s,p:=0,1
    for x>0{
        c:= x%10
        s += c
        p *= c
        x /= 10
    }
    return s, p
}