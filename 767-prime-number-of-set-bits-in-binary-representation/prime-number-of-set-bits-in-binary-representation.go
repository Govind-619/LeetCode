var primeSet = map[int]bool{
    2:true, 3:true, 5:true, 7:true, 11:true, 13:true,
    17:true, 19:true, 23:true, 29:true, 31:true,
}

func countPrimeSetBits(left int, right int) int {
    count := 0
    for i := left; i <= right; i++ {
        if primeSet[countOnes(i)] {
            count++
        }
    }
    return count
}

func countOnes(x int) int {
    count := 0
    for x > 0 {
        x = x & (x - 1)
        count++
    }
    return count
}
