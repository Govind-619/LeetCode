func reverseDegree(s string) int {
    r:= []rune(s)
    sum:=0
    for i,v:= range r{
        k:= int(v)-96
        product:= (27-k)*(i+1)
        sum += product
    }
    return sum
}