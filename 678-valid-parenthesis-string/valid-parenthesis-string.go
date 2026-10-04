func checkValidString(s string) bool {
    low, high := 0, 0
    for _, ch := range s {
        if ch == '(' {
            low++
            high++
        } else if ch == ')' {
            low--
            high--
        } else if ch == '*' {
            low--     // treat '*' as ')'
            high++    // treat '*' as '('
        }

        if high < 0 {
            return false // too many ')'
        }

        if low < 0 {
            low = 0 // can't go below 0
        }
    }
    return low == 0
}
