func countPartitions(nums []int) int {
    c:=0
    for i:=1; i<len(nums); i++{
        a:=sum(nums[:i])
        b:=sum(nums[i:])
        s:= a-b
        if s%2==0{
            c++
        }
    }
    return c
}

func sum(a []int)int{
    s:=0
    for _,v:= range a {
        s+=v
    }
    return s
}