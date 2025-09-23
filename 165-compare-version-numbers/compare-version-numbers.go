func compareVersion(version1 string, version2 string) int {
    s1 := strings.Split(version1, ".")
    s2 := strings.Split(version2, ".")
    n := max(len(s1), len(s2))

    for i := 0; i < n; i++ {
        var v1, v2 int
        if i < len(s1) {
            v1, _ = strconv.Atoi(s1[i])
        }
        if i < len(s2) {
            v2, _ = strconv.Atoi(s2[i])
        }

        if v1 < v2 {
            return -1
        } else if v1 > v2 {
            return 1
        }
    }
    return 0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
