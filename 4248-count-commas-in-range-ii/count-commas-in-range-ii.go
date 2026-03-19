func countCommas(n int64) int64 {
    var commas int64 = 0
    var start int64 = 1
    var commasPerNumber int64 = 0

    for start <= n {
        end := start*1000 - 1
        if end > n {
            end = n
        }

        count := end - start + 1
        commas += count * commasPerNumber

        commasPerNumber++
        start *= 1000
    }

    return commas
}