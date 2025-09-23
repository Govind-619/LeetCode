func smallestNumber(n int, t int) int {
    for i:= n; i<=100; i++{
        if p(i)%t==0{
            return i
        }
    }
    return 0
}

func p (x int)int{
    p:= 1
    for x>0{
        p *= x%10
        x /= 10
    }
    return p
}