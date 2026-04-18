func mirrorDistance(n int) int {
    return abs(reverse(n)-n)
}

func reverse(x int)int{
    s:=0
    for x>0{
        s=s*10 + x%10
        x/=10
    }
    return s
}

func abs(x int)int{
    if x<0{
        return -x
    }
    return x
}