func minimumBoxes(apple []int, capacity []int) int {
    s:= 0
    for _,v:= range apple{
        s+=v
    } 
    count:= 0
    sort.Ints(capacity)
    for i:= len(capacity)-1; i>=0; i--{
        if s<=0{
            return count
        }
        count++
        s-=capacity[i]
    }
    return len(capacity)
}