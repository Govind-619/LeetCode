func specialTriplets(nums []int) int {
    ml := make(map[int]int)
    mr := make(map[int]int)
    var result int
    const MOD int = 1000000007
    for _,v := range nums {
        mr[v]++
    }
    for i,v := range nums {
        if i > 0 {
            ml[nums[i-1]]++
        }
        mr[v]--
        result = (result + (mr[v*2]*ml[v*2])) % MOD
    }
    return result
}