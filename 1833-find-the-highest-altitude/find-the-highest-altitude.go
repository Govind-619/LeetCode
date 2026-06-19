func largestAltitude(gain []int) int {
    max,m:=0,0
    for _,v:= range gain{
        m+=v
        if m>max{
            max= m
        }
    }
    return max
}