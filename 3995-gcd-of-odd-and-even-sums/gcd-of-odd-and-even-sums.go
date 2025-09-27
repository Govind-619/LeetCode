func gcdOfOddEvenSums(n int) int {
    sumOdd:=n*n
    sumEven:= n*(n+1) 
    return gcd(sumOdd, sumEven)  
}

func gcd(x,y int)int{
    for y!=0{
        x,y = y,x%y
    }
    return x
}