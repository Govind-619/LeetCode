func findThePrefixCommonArray(A []int, B []int) []int {
    n := len(A)
    C := make([]int, n)
    frequency := make([]int, n+1)
    commonCount := 0

    for i := 0; i < n; i++ {
        frequency[A[i]]++
        if frequency[A[i]] == 2 {
            commonCount++
        }

        frequency[B[i]]++
        if frequency[B[i]] == 2 {
            commonCount++
        }

        C[i] = commonCount
    }

    return C
}