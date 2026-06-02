func earliestFinishTime(lst []int, ldu []int, wst []int, wdu []int) int {
    ans := int(1e9)

    for i := 0; i < len(lst); i++ {
        t := lst[i] + ldu[i]

        for j := 0; j < len(wst); j++ {
            if max(t, wst[j])+wdu[j] < ans {
                ans = max(t, wst[j]) + wdu[j]
            }
        }
    }

    for i := 0; i < len(wst); i++ {
        t := wst[i] + wdu[i]

        for j := 0; j < len(lst); j++ {
            if max(t, lst[j])+ldu[j] < ans {
                ans = max(t, lst[j]) + ldu[j]
            }
        }
    }

    return ans
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}