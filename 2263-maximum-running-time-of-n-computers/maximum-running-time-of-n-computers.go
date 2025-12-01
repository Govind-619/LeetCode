func maxRunTime(n int, batteries []int) int64 {
    q:= len(batteries)
    sort.Ints(batteries)
    live := make([]int64, n)
    for i := 0; i < n; i++{
        live[i] = int64(batteries[q - n + i])
    }
    extra := int64(0)
    for i := 0; i < q - n; i++{
        extra += int64(batteries[i])
    }
    for i := 0; i < n-1; i++{
        if extra < int64(i + 1) * (live[i + 1] - live[i]){
            return live[i] + extra / int64(i + 1)
        }
        extra -= int64(i + 1) * (live[i + 1] - live[i])
    }
    return live[n - 1] + extra / int64(n)
}