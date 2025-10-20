func decimalRepresentation(n int) []int {
    arr:= []int{}
    count:=0.0
    for n>0{
        c:= n%10
        if c!=0{
            pow:=int(math.Pow(10,count))
            arr= append([]int{c*pow},arr...)
        }
        n/=10
        count++
    }
    return arr
}