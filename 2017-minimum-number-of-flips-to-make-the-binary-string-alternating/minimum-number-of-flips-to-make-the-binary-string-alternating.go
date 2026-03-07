func minFlips(s string) int {
    // concat s onto itself to see rotations
    full := s + s
    ans := len(s)

    switches := 0
    for i, c := range full {
        // cast to rune & add unicode value of char '0'
        if rune('0' + i % 2) != c {
            switches += 1
        }
        // set back to window size of len(s)
        if i >= len(s) {
            if rune('0' + (i-len(s)) % 2) != c {
                switches -= 1
            }
            // compare with current running minimum num switches needed
            ans = min(ans, switches, len(s)-switches)
        }
    }
    return ans
}