func maxSideLength(mat [][]int, threshold int) int {
    m := len(mat)    
    n := len(mat[0])
    pref := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        pref[i] = make([]int, n+1)
    }

    for i:= 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            pref[i][j] = mat[i-1][j-1] + pref[i-1][j] + pref[i][j-1] - pref[i-1][j-1]
        }
    }

    ans := 0
    lo, hi := 1, min(n, m)

    for lo <= hi {
        mid := (lo + hi) >> 1
        if check(pref, mid, threshold) {
            ans = mid
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }

    return ans
}

func check(pref [][]int, k, threshold int) bool {
    for i := k; i < len(pref); i++ {
        for j := k; j < len(pref[0]); j++ {
            sum := pref[i][j] - pref[i-k][j] - pref[i][j-k] + pref[i-k][j-k]
            if sum <= threshold {
                return true
            }
        }
    }
    return false
}