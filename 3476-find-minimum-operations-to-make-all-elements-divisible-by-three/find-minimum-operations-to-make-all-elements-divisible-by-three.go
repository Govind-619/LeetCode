func minimumOperations(nums []int) int {
    count:=0
    for _,v:=range nums{
        if v%3!=0{
            count++
        }
    }
    return count
}