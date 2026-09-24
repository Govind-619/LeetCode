func smallestIndex(nums []int) int {
    for i,v:= range nums{
        sum:=0
        for v>0{
            sum += v%10
            v /= 10
        }
        if sum == i{
            return i
        }
    }
    return -1
}