func minBitwiseArray(nums []int) []int {
    res := make([]int, len(nums))
    for i, n := range nums {
        res[i] = findNum(n)
    }
    return res
}

func findNum(x int) int {
    for i := 1; i < x; i++ {
        if i | (i+1) == x {
            return i
        }
    }
    return -1
}