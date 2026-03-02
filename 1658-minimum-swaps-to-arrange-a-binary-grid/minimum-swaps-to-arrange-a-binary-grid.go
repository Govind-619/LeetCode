func minSwaps(grid [][]int) int {
    n := len(grid)
    var rightZeroCnt []int = make([]int, 0)
    for i := range n {
        cnt := 0
        for j := n - 1; j >= 0; j-- {
            if grid[i][j] == 0 {
                cnt++
            } else {
                break
            }
        }
        rightZeroCnt = append(rightZeroCnt, cnt)
    }
    var res int = 0
    for i := 0; i < n; i++ {
        if rightZeroCnt[i] < n - 1 - i {
            j := i + 1
            for j < n && rightZeroCnt[j] < n - 1 - i {
                j++
            }
            if j == n {
                return -1
            }
            temp := rightZeroCnt[j]
            res += j - i
            for j > i {
                rightZeroCnt[j] = rightZeroCnt[j - 1]
                j--
            }
            rightZeroCnt[i] = temp
        }
    }
    return res
}