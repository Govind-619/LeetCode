func maxProduct(n int) int {
    res:= []int{}
    for n>0{
        r:= n%10
        res= append(res,r)
        n /= 10
    }
    sort.Ints(res)
    return res[len(res)-1]*res[len(res)-2]
}