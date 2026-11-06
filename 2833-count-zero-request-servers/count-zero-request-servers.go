func countServers(n int, logs [][]int, x int, queries []int) []int {
    slices.SortFunc(logs, func(a, b []int) int { return cmp.Compare(a[1], b[1]) })  // by time
    logs = append(logs, []int{0, math.MaxInt})  // sentinel: no window ever reaches it
    order := make([]int, len(queries))
    for i := range order {
        order[i] = i
    }
    slices.SortFunc(order, func(i, j int) int { return cmp.Compare(queries[i], queries[j]) })  // query positions in time order

    cnt := make([]int, n+1)  // requests per server inside the window [q - x, q]
    ans := make([]int, len(queries))
    busy, lo, hi := 0, 0, 0  // busy: servers with at least one request in it
    for _, i := range order {
        q := queries[i]
        for ; logs[hi][1] <= q; hi++ {  // requests up to q come in
            if cnt[logs[hi][0]] == 0 {
                busy++
            }
            cnt[logs[hi][0]]++
        }
        for ; logs[lo][1] < q-x; lo++ {  // requests before q - x go out
            cnt[logs[lo][0]]--
            if cnt[logs[lo][0]] == 0 {
                busy--
            }
        }
        ans[i] = n - busy
    }
    return ans
}