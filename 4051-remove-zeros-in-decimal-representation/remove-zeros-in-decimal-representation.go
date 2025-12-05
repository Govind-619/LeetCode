func removeZeros(n int64) int64 {
    no:=""
    for n>0{
        if n%10!=0{
            s:= strconv.Itoa(int(n%10))
            no= s+no
        }
        n/=10
    }
    c,_:= strconv.Atoi(no)
    return int64(c)
}