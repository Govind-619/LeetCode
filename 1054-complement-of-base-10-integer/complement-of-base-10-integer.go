func bitwiseComplement(n int) int {
     b:= strconv.FormatInt(int64(n), 2)
    c:=[]rune(b)
    for i,v := range c{
        if v=='0'{
            c[i]= '1'
        }else{
            c[i]= '0'
        }
    } 
    d,_:= strconv.ParseInt(string(c), 2, 64)
    return int(d)
}