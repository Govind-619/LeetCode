func minimumFlips(n int) int {
    b:= strconv.FormatInt(int64(n),2)
    r:= reverse(b)
    c:= compare(b,r)
    return c
}

func reverse(x string)string{
    r:= []rune(x)
    for i,j:=0, len(r)-1; i<j; i,j= i+1, j-1{
        r[i],r[j]  =r[j],r[i]
    }
    return string(r)
} 

func compare(x,y string)int{
    count:=0
    for i:=0; i<len(x); i++{
        if x[i]!=y[i]{
            count++
        }
    }
    return count
}