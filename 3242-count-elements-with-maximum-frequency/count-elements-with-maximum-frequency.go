func maxFrequencyElements(nums []int) int {
    m:= make(map[int]int)

    max:= 0
    for _,v:= range nums{
        m[v]++
        if m[v]>max{
            max = m[v]
        }
    }
    count:=0
    for _, val:= range m{
        if val == max{
            count += val
        }
    }
    
    return count
}