func sortByBits(arr []int) []int {
    sort.Slice(arr, func(m, n int) bool {
        a := bits.OnesCount(uint(arr[m]))
        b := bits.OnesCount(uint(arr[n]))
        return (a == b && arr[m] < arr[n]) || a < b
    })

    return arr
}