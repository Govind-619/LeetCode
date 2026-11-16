func digitSum(s string, k int) string {
    for len(s) > k {
        var next strings.Builder
        for i := 0; i < len(s); i += k {
            sum := 0
            for _, d := range s[i:min(i+k, len(s))] {  // one group; the last may be shorter
                sum += int(d - '0')
            }
            next.WriteString(strconv.Itoa(sum))
        }
        s = next.String()
    }
    return s
}