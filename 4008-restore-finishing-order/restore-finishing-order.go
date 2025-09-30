func recoverOrder(order []int, friends []int) []int {
    var (
        res = make([]int, len(friends))
        i int
    )

    for _, v := range order {
        if slices.Contains(friends, v) {
            res[i] = v
            i++
        }
    }

    return res
}