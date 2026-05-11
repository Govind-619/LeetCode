func separateDigits(nums []int) []int {
    res:=[]int{}
    for _,v:= range nums{
        a:=[]int{}
        for v>9{
            a=append([]int{v%10},a...)
            v/=10
        }
        a=append([]int{v},a...)
        res= append(res,a...)
    }
    return res
}