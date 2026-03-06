func checkOnesSegment(s string) bool {
    cnt := 0

    if s[0] == '1' {
        cnt++
    }

    for i := 0; i+1 < len(s); i++ {
        if s[i] == '0' && s[i+1] == '1' {
            cnt++
        }
    }

    return cnt <= 1
}