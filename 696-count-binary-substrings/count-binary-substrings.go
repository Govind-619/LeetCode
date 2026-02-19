func countBinarySubstrings(s string) int {
    ans := 0
    cur := s[0]
    countCur := 1
    countNotCur := 0
    for i := 1; i < len(s); i++ {
        if cur == s[i] {
            countCur++
        } else {
            cur = s[i]
            countNotCur = countCur
            countCur = 1
        }
        if countCur <= countNotCur {
            ans++
        }
    }
    return ans
}