func maximumTotalDamage(power []int) int64 {
    freq := make(map[int]int64)
    for _, p := range power {
        freq[p]++
    }

    arr := make([]int, 0, len(freq))
    for k, _ := range freq {
        arr = append(arr, k)
    }
    sort.Ints(arr)

    n := len(arr)
    dp := make([]int64, n)
    dp[0] = freq[arr[0]] * int64(arr[0])
    for i := 1; i < n; i++ {
        dmg := freq[arr[i]] * int64(arr[i])
        prev := sort.Search(i, func(k int) bool {
            return arr[k] > arr[i] - 3
        }) - 1
        if prev >= 0 {
            dmg += dp[prev]
        }
        dp[i] = max(dp[i-1], dmg)
    }

    return dp[n-1]
}