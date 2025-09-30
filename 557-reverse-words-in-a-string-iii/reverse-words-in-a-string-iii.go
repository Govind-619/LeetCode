func reverseWords(s string) string {
    c:= strings.Fields(s)
    for i,v:= range c{
        c[i]= reverse(v)
    }
    str:= strings.Join(c," ")
    return str
}

func reverse(a string)string{
    r:= []rune(a)
    for i,j:=0, len(r)-1; i<j; i,j = i+1, j-1{
        r[i], r[j]= r[j], r[i]
    }
    return string(r)
}