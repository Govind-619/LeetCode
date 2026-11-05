func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
    slices.Sort(arr2)
    arr2 = append(arr2, math.MaxInt)  // sentinel: some b is always >= a - d
    count := 0
    for _, a := range arr1 {
        i, _ := slices.BinarySearch(arr2, a-d)  // the first b >= a - d
        if arr2[i] > a+d {
            count++
        }
    }
    return count
}