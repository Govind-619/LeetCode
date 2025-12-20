func minDeletionSize(strs []string) int {
    delete:= 0
    for col:=0; col<len(strs[0]); col++{
        for row:=0; row<len(strs)-1; row++{
            if strs[row][col]>strs[row+1][col]{
                delete++
                break
            }
        }

    }
    return delete
}