func maxFrequencyElements(nums []int) int {
    m:= make(map[int]int)

    max:= 0
    count:=0
    for _,v:= range nums{
        m[v]++
        if m[v]>max{
            max = m[v]
            count=1
        }else if m[v]== max{
            count++
        }
    }
    
   
    return count* max
}