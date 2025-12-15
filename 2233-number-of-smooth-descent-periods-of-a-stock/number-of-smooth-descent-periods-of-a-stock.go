func getDescentPeriods(prices []int) int64 {
    var ans int64 = 1
    var currLen int64 = 1

    for i := 1; i < len(prices); i++ {
        if prices[i] == prices[i-1]-1 {
            currLen++
        } else {
            currLen = 1
        }
        ans += currLen
    }

    return ans
}
