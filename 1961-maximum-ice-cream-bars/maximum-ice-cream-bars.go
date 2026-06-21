func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)
    count:=0
    for _,v:= range costs{
        if coins<=0 ||coins<v{
            break
        }
        coins -= v
        count++ 
    }
    return count
}