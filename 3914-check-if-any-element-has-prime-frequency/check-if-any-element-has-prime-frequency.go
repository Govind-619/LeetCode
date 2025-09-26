func checkPrimeFrequency(nums []int) bool {
    m:= make(map[int]int)
    for _,v:= range nums{
        m[v]++
    }

    for _,v:= range m{
        if prime(v){
            return true
        }
    }
    return false
}

func prime(x int)bool{
    if x<2{
        return false
    }
    for i:=2; i*i<=x; i++{
        if x%i==0{
            return false
        }
    }
    return true
}