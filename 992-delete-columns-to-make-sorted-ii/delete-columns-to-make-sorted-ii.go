func minDeletionSize(strs []string) int {
    count := 0
    increasing := make([]bool, len(strs))
    increasingColumn := make([]bool, len(strs))
    for i := 0; i < len(strs[0]); i++ {
        toDelete := false
        copy(increasingColumn, increasing)
        for j := 1; j < len(strs); j++ {
            if increasingColumn[j] {
                continue
            }
            if strs[j-1][i] > strs[j][i] {
                toDelete = true
                break
            } else if strs[j-1][i] < strs[j][i] {
                increasingColumn[j] = true
            }
        }
        if toDelete {
            count++
            continue
        }
        copy(increasing, increasingColumn)
    }
    return count
}